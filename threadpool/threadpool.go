package threadpool

import (
	"time"
)

const (
	MAX_THREAD_COUNT = 102400
)

type ThreadPool struct {
	taskList     chan ITask
	runningCount chan bool
}

type ITask interface {
	onExecute()
}

func (self *ThreadPool) initThreads() {
	for i := 0; i < MAX_THREAD_COUNT; i++ {
		go func() {
			task := <-self.taskList
			if task != nil {
				task.onExecute()
			}
			<-self.runningCount
		}()
	}
}

func (self *ThreadPool) Start(task ITask) {
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
			print("wait", len(self.runningCount))
			if len(self.runningCount) == 0 {
				return
			}
		}
	}
}

func New() *ThreadPool {
	object := &ThreadPool{
		taskList:     make(chan ITask, MAX_THREAD_COUNT),
		runningCount: make(chan bool, MAX_THREAD_COUNT),
	}

	object.initThreads()

	return object
}
