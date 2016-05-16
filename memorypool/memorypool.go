package memorypool

import (
	"unsafe"
)

const (
	MAX_MEM_LIST_SIZE = 10240
)

var instance *MemoryPool

type MemoryPool struct {
	data map[uint]chan []byte
	lock chan bool
}

func New() *MemoryPool {
	pool := new(MemoryPool)
	pool.data = make(map[uint]chan []byte)
	pool.lock = make(chan bool, 1)

	return pool
}

func (self *MemoryPool) Alloc(size uint) ([]byte, bool) {
	self.lock <- true
	defer func() { <-self.lock }()

	if len(self.data[size]) > 0 {
		return <-self.data[size], false
	}

	buffer := make([]byte, size+4)
	ptr := (*uint)(unsafe.Pointer(&buffer[0]))
	*ptr = size

	return buffer[4:], true
}

func (self *MemoryPool) Free(buffer []byte) {
	self.lock <- true
	defer func() { <-self.lock }()

	ptr := (*uint)(unsafe.Pointer(uintptr(unsafe.Pointer(&buffer[0])) - uintptr(4)))
	size := *ptr

	data := self.data
	if data[size] == nil {
		data[size] = make(chan []byte, MAX_MEM_LIST_SIZE)
	}

	self.data[size] <- buffer
}

func (self *MemoryPool) Clean() {
	self.lock <- true
	defer func() { <-self.lock }()

	for _, ch := range self.data {
		close(ch)
	}

	self.data = make(map[uint]chan []byte)
}

func GetInstance() *MemoryPool {
	if instance == nil {
		instance = New()
	}

	return instance
}
