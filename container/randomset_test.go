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
		t.Fatal("1")
	}

	ok = set.Remove(1)
	if ok {
		t.Fatal("2")
	}

	set.Insert(1)
	set.Insert(2)
	set.Insert(3)

	ok = set.Insert(1)
	if ok {
		t.Fatal("3")
	}

	ok = set.Remove(1)
	if !ok {
		t.Fatal("4")
	}

	ok = set.Remove(1)
	if ok {
		t.Fatal("5")
	}

	ok = set.Has(2)
	if !ok {
		t.Fatal("6")
	}

	set.Remove(2)

	if _, ok := set.RandomAndSkip(3); ok {
		t.Fatal("7")
	}

	set.Remove(3)
	if _, ok := set.Random(); ok {
		t.Fatal("8")
	}

	set.Reset()
	if set.Len() != 0 {
		t.Fatal("9")
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
			t.Fatal("10")
		}
		data[v.(int)] = true
	}

	for i := 0; i < 5; i++ {
		if !data[i] {
			t.Fatal("11")
		}
	}

	data = make(map[int]bool)
	for i := 0; i < 100; i++ {
		v, ok := set.RandomAndSkip(3)
		if !ok {
			t.Fatal("12")
		}
		data[v.(int)] = true
	}

	if data[3] {
		t.Fatal("13")
	}

	if set.Len() != 5 {
		t.Fatal("14")
	}

	list, _ := set.GetList(5)
	if len(list) != 5 {
		t.Fatal("15")
	}

	data = make(map[int]bool)
	for _, v := range list {
		data[v.(int)] = true
	}

	for i := 0; i < 5; i++ {
		if !data[i] {
			t.Fatal("16")
		}
	}

	// for i := 0; i < 100; i++ {
	// 	list, _ := set.GetListAndSkip(3, 10)
	// 	for _, v := range list {
	// 		log.Println(v)
	// 	}
	// 	if len(list) != 4 {
	// 		t.Fatal("17")
	// 	}

	// 	data = make(map[int]bool)
	// 	for _, v := range list {
	// 		data[v.(int)] = true
	// 	}

	// 	// for j := 0; j < 5; j++ {
	// 	// 	if !data[j] {
	// 	// 		t.Fatal("18")
	// 	// 	}
	// 	// }

	// 	// if data[3] {
	// 	// 	t.Fatal("19")
	// 	// }
	// }

	if set.GetFreeCount() != 95 {
		t.Fatal("free count:", set.GetFreeCount())
	}

	println(set.GetAll())
}
