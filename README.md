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
