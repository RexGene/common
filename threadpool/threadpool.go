package threadpool

import (
	"time"
)

const (
	MAX_THREAD_COUNT = 8
)

var instance *ThreadPool

type ThreadPool struct {
	taskList chan func()
}

func (self *ThreadPool) initThreads() {
	for i := 0; i < MAX_THREAD_COUNT; i++ {
		go func() {
			for {
				task := <-self.taskList
				if task != nil {
					task()
				}
			}
		}()
	}
}

func (self *ThreadPool) Start(task func()) {
	self.taskList <- task
}

func (self *ThreadPool) Stop() {
	close(self.taskList)
}

func (self *ThreadPool) WaitAllFinish() {
	for {
		select {
		case <-time.After(time.Second):
			if len(self.taskList) == 0 {
				return
			}
		}
	}
}

func New() *ThreadPool {
	object := &ThreadPool{
		taskList: make(chan func(), MAX_THREAD_COUNT),
	}

	object.initThreads()

	return object
}

func GetInstance() *ThreadPool {
	if instance == nil {
		instance = New()
	}

	return instance
}
