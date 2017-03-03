package timingwheel

import (
	"container/list"
	"container/ring"
)

type BaseNode struct {
	isEnable   bool
	isRemove   bool
	handlecall func()
}

type TimingWheel struct {
	r *ring.Ring
}

func Remove(node *BaseNode) {
	node.isRemove = true
	node.isEnable = false
}

func New(amount uint) *TimingWheel {
	object := &TimingWheel{
		r: ring.New(int(amount)),
	}

	return object
}

func (self *TimingWheel) Tick() {
	self.r = self.r.Next()
	var element *list.List
	if nil == self.r.Value {
		element = list.New()
		self.r.Value = element
	} else {
		element = self.r.Value.(*list.List)
	}

	iter := element.Front()
	for iter != nil && iter.Value != nil {
		base := iter.Value.(*BaseNode)
		if base.isEnable {
			base.handlecall()
		}

		nextIter := iter.Next()
		if base.isRemove {
			element.Remove(iter)
		}
		iter = nextIter
	}
}

func (self *TimingWheel) InsertCallback(handlecall func()) *BaseNode {
	node := &BaseNode{
		isEnable:   true,
		isRemove:   true,
		handlecall: handlecall,
	}

	var current *list.List
	if nil == self.r.Value {
		current = list.New()
		self.r.Value = current
	} else {
		current = self.r.Value.(*list.List)
	}

	current.PushBack(node)

	return node
}

func (self *TimingWheel) InsertCallbackForever(handlecall func()) *BaseNode {
	node := &BaseNode{
		isEnable:   true,
		isRemove:   false,
		handlecall: handlecall,
	}

	var current *list.List
	if nil == self.r.Value {
		current = list.New()
		self.r.Value = current
	} else {
		current = self.r.Value.(*list.List)
	}

	current.PushBack(node)

	return node
}
