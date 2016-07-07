package timingwheel

import (
	"testing"
)

var call_count = 0

func test() {
	println("I'm hit")
	call_count++
}

func TestInsert(t *testing.T) {
	call_count = 0

	object := New(1)
	object.InsertCallback(test)

	for i := 0; i < 10; i++ {
		object.Tick()
	}

	if call_count != 1 {
		t.Log("call count error!")
		t.Fatal()
	}
}

func TestRemove(t *testing.T) {
	call_count = 0

	object := New(2)
	handler := object.InsertCallbackForever(test)

	for i := 0; i < 10; i++ {
		object.Tick()
	}

	if call_count != 5 {
		t.Log("forever call count error!")
		t.Fatal()
	}

	call_count = 0
	object.Remove(handler)

	for i := 0; i < 2; i++ {
		object.Tick()
	}

	if call_count != 0 {
		t.Log("after remove call count error!")
		t.Fatal()
	}
}
