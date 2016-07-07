# golang common library

## MemoryPool 

### Install

```
go get github.com/RexGene/common/memorypool
```

### Usage

```golang

import "github.com/RexGene/common/memorypool"

mempool = memorypool.GetInstance() 
buff := mempool.Alloc(1000)
mempool.Free(buff)

mempool.Clean()


```

## ThreadPool

### Install
```
go get github.com/RexGene/common/threadpool
```
### Usage

```golang

import "github.com/RexGene/common/threadpool"

type TestTask struct {
}

func (self TestTask) onExecute() {
	print("test")
}

pool := threadpool.GetInstance()
task := new(TestTask)
for i := 0; i < 1000; i++ {
    pool.Start(task.onExecute)
}

pool.WaitAllFinish()


```
## ObjectPool


### Install
```
go get github.com/RexGene/common/objectpool
```
### Usage

```golang

import "github.com/RexGene/common/objectpool"

type TestObject struct {
    objectpool.BaseObject
}

func TextNew() IPoolObject {
	return &TestObject{}
}

pool := objectpool.GetInstance()
pool.RegistObject(1, TextNew)

object := pool.MakeObject(1)
pool.RecoverObject(object)

```

### Install
```
go get github.com/RexGene/common/timingwheel
```

### Usage
```golang
import (
    "time"
    "github.com/RexGene/common/timingwheel"
)

tw := timingwheel.New(60)
var callback = func() {
    println("call in")
}

tw.InsertCallback(callback)

for {
    select {
        case <-time.After(time.Second):
            tw.Tick()
    }
}

```
