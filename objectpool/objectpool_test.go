package objectpool

import (
	"testing"
)

type TestObject struct {
	BaseObject
}

func TextNew() IPoolObject {
	return &TestObject{}
}

func TestObjectPool(t *testing.T) {
	pool := GetInstance()

	pool.RegistObject(1, TextNew)
	object, isNew := pool.MakeObject(1)
	if !isNew {
		t.Log("first make object must be new")
		t.Fail()
		return
	}

	println("object id:", object.GetObjectType())
	pool.RecoverObject(object)
	object, isNew = pool.MakeObject(1)
	if isNew {
		t.Log("first make object must be not new")
		t.Fail()
		return
	}
}
