package threadpool

import (
	"time"
)

const (
	MAX_THREAD_COUNT = 10240
)

var instance *ThreadPool

type ThreadPool struct {
	taskList     chan func()
	runningCount chan bool
}

func (self *ThreadPool) initThreads() {
	for i := 0; i < MAX_THREAD_COUNT; i++ {
		go func() {
			task := <-self.taskList
			if task != nil {
				task()
			}
			<-self.runningCount
		}()
	}
}

func (self *ThreadPool) Start(task func()) {
	self.taskList <- task
	self.runningCount <- true
}

func (self *ThreadPool) Stop() {
	close(self.taskList)
	close(self.runningCount)
}

func (self *ThreadPool) WaitAllFinish() {
	for {
		select {
		case <-time.After(time.Second):
			if len(self.runningCount) == 0 {
				return
			}
		}
	}
}

func New() *ThreadPool {
	object := &ThreadPool{
		taskList:     make(chan func(), MAX_THREAD_COUNT),
		runningCount: make(chan bool, MAX_THREAD_COUNT),
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
