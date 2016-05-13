# golang common library

## MemoryPool 

### Install

```
go get github.com/RexGene/common/memorypool
```

### Usage

```golang

import "github.com/RexGene/common/memorypool"

mempool = memorypool.New() 
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

pool := threadpool.New()
task := new(TestTask)
for i := 0; i < 1000; i++ {
    pool.Start(task)
}

pool.WaitAllFinish()


```

