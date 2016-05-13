package memorypool

import (
	"testing"
)

func TestAllocAndFree(t *testing.T) {
	var mempool = New()

	for i := 1; i < 10000; i++ {
		buffer, isNew := mempool.Alloc(uint(i))
		mempool.Free(buffer)

		if !isNew {
			t.Log("first create must new")
			t.Fail()
			return
		}
	}

	mempool.Clean()
}

func TestDoubleAllocAndFree(t *testing.T) {
	var mempool = New()
	isFirst := true

	for i := 1; i < 10000; i++ {
		buffer, isNew := mempool.Alloc(uint(1000))
		buffer1, isNew1 := mempool.Alloc(uint(1000))
		mempool.Free(buffer)
		mempool.Free(buffer1)

		if isFirst {
			if !isNew || !isNew1 {
				t.Log("first create must new")
				t.Fail()
				return
			}

			isFirst = false
		} else {
			if isNew || isNew1 {
				t.Log("second create must not new")
				t.Fail()
				return
			}
		}

	}

	mempool.Clean()
}
