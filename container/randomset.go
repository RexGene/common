package container

import (
	"math/rand"
)

const (
	DEFAULT_CAP = 128
)

type Value interface{}

type RandomSet struct {
	dataMap    map[Value]uint
	dataList   []Value
	randomList []Value
	freeCount  uint
}

func NewRandomSet() *RandomSet {
	valueList := make([]Value, 0, DEFAULT_CAP)
	return &RandomSet{
		dataMap:    make(map[Value]uint),
		dataList:   valueList,
		freeCount:  0,
		randomList: valueList,
	}
}

func (self *RandomSet) Insert(v Value) bool {
	_, ok := self.dataMap[v]
	if ok {
		return false
	}

	freeIndex := uint(0)
	if self.freeCount > 0 {
		freeIndex = self.freeCount - 1
		self.freeCount--

		self.dataList[freeIndex] = v
		self.randomList = self.dataList[self.freeCount:]
	} else {
		freeIndex = uint(len(self.dataList))
		self.dataList = append(self.dataList, v)
		self.randomList = self.dataList[self.freeCount:]
	}

	self.dataMap[v] = freeIndex
	return true
}

func (self *RandomSet) Remove(v Value) bool {
	freeCount := self.freeCount
	if freeCount >= uint(len(self.dataList)) {
		return false
	}

	index, ok := self.dataMap[v]
	if !ok {
		return false
	}

	dataList := self.dataList

	temp := dataList[freeCount]
	dataList[freeCount] = dataList[index]
	dataList[index] = temp

	self.dataMap[temp] = index
	delete(self.dataMap, v)

	self.freeCount++
	self.randomList = dataList[self.freeCount:]

	return true
}

func (self *RandomSet) Has(v Value) bool {
	_, ok := self.dataMap[v]
	return ok
}

func (self *RandomSet) RandomAndSkip(skip Value) (Value, bool) {
	randomSize := len(self.randomList)
	if randomSize <= 1 {
		return nil, false
	}

	randValue := uint(rand.Int31n(int32(randomSize) - 1))
	index, ok := self.dataMap[skip]
	if ok {
		if self.freeCount+randValue >= index {
			randValue++
		}
	}

	return self.randomList[randValue], true
}

func (self *RandomSet) Random() (Value, bool) {
	randomSize := len(self.randomList)
	if randomSize <= 0 {
		return nil, false
	}

	randValue := uint(rand.Int31n(int32(randomSize)))
	return self.randomList[randValue], true
}

func (self *RandomSet) Len() int {
	return len(self.randomList)
}

func (self *RandomSet) Reset() {
	self.dataMap = make(map[Value]uint)
	self.dataList = make([]Value, 0, DEFAULT_CAP)
	self.randomList = self.dataList
	self.freeCount = 0
}

func (self *RandomSet) GetFreeCount() uint {
	return self.freeCount
}
