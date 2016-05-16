package objectpool

type BaseObject struct {
	typeId int
}

func (self *BaseObject) GetObjectType() int {
	return self.typeId
}

func (self *BaseObject) SetObjectType(typeId int) {
	self.typeId = typeId
}
