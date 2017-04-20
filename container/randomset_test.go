package container

import (
	"testing"
)

func TestRandom(t *testing.T) {
	set := NewRandomSet()
	set.Insert(1)
	set.Insert(2)
	set.Insert(3)

	ok := set.Insert(1)
	if ok {
		t.Fatal("insert has a error")
	}

	ok = set.Remove(1)
	if !ok {
		t.Fatal("Remove has a error")
	}

	ok = set.Remove(1)
	if ok {
		t.Fatal("Remove has a error")
	}

	ok = set.Has(2)
	if !ok {
		t.Fatal("Has function has a error")
	}

	set.Remove(2)

	if _, ok := set.RandomAndSkip(3); ok {
		t.Fatal("random function has a error")
	}

	set.Remove(3)
	if v, ok := set.Random(); ok {
		t.Fatal("random function has a error:", v)
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
			t.Fatal("random has a error")
		}
		data[v.(int)] = true
	}

	for i := 0; i < 5; i++ {
		if !data[i] {
			t.Fatal("random function has a error")
		}
	}

	data = make(map[int]bool)
	for i := 0; i < 100; i++ {
		v, ok := set.RandomAndSkip(3)
		if !ok {
			t.Fatal("random has a error")
		}
		data[v.(int)] = true
	}

	if data[3] {
		t.Fatal("RandomAndSkip has a error")
	}

	if set.Len() != 5 {
		t.Fatal("len has a error")
	}
}
