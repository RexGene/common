package memorypool

import (
	"sync"
)

const (
	MAX_MEM_LIST_SIZE = 10240
)

var instance *MemoryPool

type MemoryPool struct {
	sync.Mutex
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
	self.Lock()
	defer self.Unlock()

	if len(self.data[size]) > 0 {
		buffer := <-self.data[size]
		println("alloc ex:", buffer, " size:", size)
		return buffer, false
	}

	buffer := make([]byte, size)

	println("alloc:", buffer, " size:", size)
	return buffer[:], true
}

func (self *MemoryPool) Free(buffer []byte) {
	self.Lock()
	defer self.Unlock()

	size := uint(cap(buffer))

	data := self.data
	if data[size] == nil {
		data[size] = make(chan []byte, MAX_MEM_LIST_SIZE)
	}

	println("free:", buffer, " size:", size)
	self.data[size] <- buffer
}

func (self *MemoryPool) Clean() {
	self.Lock()
	defer self.Unlock()

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
