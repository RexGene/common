package container

import (
	"testing"
)

func TestRandom(t *testing.T) {
	set := NewRandomSet()
	_, ok := set.RandomAndSkip(3)
	if ok {
		t.Fail()
	}

	_, ok = set.Random()
	if ok {
		t.Fail()
	}

	ok = set.Remove(1)
	if ok {
		t.Fail()
	}

	set.Insert(1)
	set.Insert(2)
	set.Insert(3)

	ok = set.Insert(1)
	if ok {
		t.Fail()
	}

	ok = set.Remove(1)
	if !ok {
		t.Fail()
	}

	ok = set.Remove(1)
	if ok {
		t.Fail()
	}

	ok = set.Has(2)
	if !ok {
		t.Fail()
	}

	set.Remove(2)

	if _, ok := set.RandomAndSkip(3); ok {
		t.Fail()
	}

	set.Remove(3)
	if _, ok := set.Random(); ok {
		t.Fail()
	}

	set.Reset()
	if set.Len() != 0 {
		t.Fail()
	}

	for i := 0; i < 100; i++ {
		set.Insert(i)
	}

	for i := 0; i < 100; i++ {
		set.Remove(i)
	}

	for i := 0; i < 5; i++ {
		set.Insert(i)
	}

	data := make(map[int]bool)
	for i := 0; i < 100; i++ {
		v, ok := set.Random()
		if !ok {
			t.Fail()
		}
		data[v.(int)] = true
	}

	for i := 0; i < 5; i++ {
		if !data[i] {
			t.Fail()
		}
	}

	data = make(map[int]bool)
	for i := 0; i < 100; i++ {
		v, ok := set.RandomAndSkip(3)
		if !ok {
			t.Fail()
		}
		data[v.(int)] = true
	}

	if data[3] {
		t.Fail()
	}

	if set.Len() != 5 {
		t.Fail()
	}
}
