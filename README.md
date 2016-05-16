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

pool := objectpool.New()
pool.RegistObject(1, TextNew)

object := pool.MakeObject(1)
pool.RecoverObject(object)

```
