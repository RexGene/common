package threadpoll

import (
	"time"
)

const (
	MAX_THREAD_COUNT = 102400
)

type ThreadPoll struct {
	taskList     chan ITask
	runningCount chan bool
}

type ITask interface {
	onExecute()
}

func (self *ThreadPoll) initThreads() {
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

func (self *ThreadPoll) Start(task ITask) {
	self.taskList <- task
	self.runningCount <- true
}

func (self *ThreadPoll) Stop() {
	close(self.taskList)
	close(self.runningCount)
}

func (self *ThreadPoll) WaitAllFinish() {
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

func New() *ThreadPoll {
	object := &ThreadPoll{
		taskList:     make(chan ITask, MAX_THREAD_COUNT),
		runningCount: make(chan bool, MAX_THREAD_COUNT),
	}

	object.initThreads()

	return object
}
