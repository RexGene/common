package objectpool

const OBJECT_LIST_SIZE = 256

type IPoolObject interface {
	GetObjectType() int
	SetObjectType(int)
}

type ObjectFunc struct {
	MakeFunc func() IPoolObject
}

type ObjectPool struct {
	pool    map[int][]IPoolObject
	funcMap map[int]ObjectFunc
}

func (self *ObjectPool) RegistObject(id int, makeFunc func() IPoolObject) {
	objectFunc := ObjectFunc{
		MakeFunc: makeFunc,
	}

	self.funcMap[id] = objectFunc
}

func (self *ObjectPool) MakeObject(id int) (IPoolObject, bool) {
	objectList := self.pool[id]
	if objectList != nil && len(objectList) != 0 {
		object := objectList[0]
		self.pool[id] = objectList[1:]
		return object, false
	}

	objectFunc := self.funcMap[id]
	if objectFunc.MakeFunc != nil {
		object := objectFunc.MakeFunc()
		object.SetObjectType(id)
		return object, true
	}

	return nil, false
}

func (self *ObjectPool) RecoverObject(object IPoolObject) {
	objectType := object.GetObjectType()
	objectList := self.pool[objectType]
	if objectList == nil {
		objectList = make([]IPoolObject, 0, OBJECT_LIST_SIZE)
	}

	objectList = append(objectList, object)
	self.pool[objectType] = objectList
}

func New() *ObjectPool {
	return &ObjectPool{
		pool:    make(map[int][]IPoolObject),
		funcMap: make(map[int]ObjectFunc),
	}
}
