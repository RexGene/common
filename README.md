# golang common library

## MemoryPool 

###Usage

```

import "github.com/RexGene/common/memorypool"

mempool = memorypool.New() 
buff := mempool.Alloc(1000)
mempool.Free(buff)

mempool.Clean()


```
