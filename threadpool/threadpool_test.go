package threadpool

import (
	"testing"
)

type TestTask struct {
}

func (self TestTask) onExecute() {
	print("test")
}

func TestStart(t *testing.T) {
	pool := New()
	task := new(TestTask)
	for i := 0; i < 1000; i++ {
		pool.Start(task.onExecute)
	}

	pool.WaitAllFinish()
}
