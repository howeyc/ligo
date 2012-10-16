// Copyright 2012 Chris Howey. All right reserved.
// Use of this source code is governed by the ISC
// license that can be found in the LICENSE file.

package ligo

import (
	"reflect"
	"testing"
)

func TestEquality(t *testing.T) {
	list1 := List(1, 2, 3, 4, 5)
	list2 := List(1, 2, 3, 4, 5)
	vec1 := Vector(1, 2, 3, 4, 5)
	vec2 := Vector(1, 2, 3, 4, 5)

	if Equal(list1, list2) == false {
		t.Fatalf("List equality failed!")
	}
	if Equal(vec1, vec2) == false {
		t.Fatalf("Vector equality failed!")
	}
	if Equal(vec1, list1) == false {
		t.Fatalf("Vector/List equality failed!")
	}
}

func TestAppend(t *testing.T) {
	list1 := List(1, 2, 3)
	list2 := List(1, 2, 3)
	listapp := List(1, 2, 3, 1, 2, 3)

	if Equal(Append(list1, list2), listapp) == false {
		t.Fatalf("Append failed!")
	}
}

func TestSubSeq(t *testing.T) {
	list1 := List(1, 2, 3, 4, 5)
	list2 := List(1, 2, 3)
	vec1 := Vector(1, 2, 3, 4, 5)
	vec2 := Vector(1, 2, 3)

	if Equal(SubSeq(list1, 0, 3), list2) == false {
		t.Fatalf("List SubSeq failed!")
	}
	if Equal(SubSeq(vec1, 0, 3), vec2) == false {
		t.Fatalf("Vector SubSeq failed!")
	}
	if Equal(SubSeq(vec1, 0, 3), Take(vec1, 3)) == false {
		t.Fatalf("Take failed!")
	}
	if Equal(SubSeq(list1, 3, 5), Drop(vec1, 3)) == false {
		t.Fatalf("Drop failed!")
	}
}

func intAdd(vals ...interface{}) interface{} {
	var result int64
	for _, val := range vals {
		v := reflect.ValueOf(val)
		intVal := v.Int()
		result += intVal
	}
	return result
}

func TestReduce(t *testing.T) {
	list1 := List(3, 4, 5)
	list2 := List(7, 5)
	result1 := Reduce(intAdd, list1)
	result2 := Reduce(intAdd, list2)
	if result1 != result2 {
		t.Fatalf("Reduce failed!, Received %v, Expected %v.", result1, result2)
	}
}

func TestMapCar(t *testing.T) {
	list1 := List(3, 4, 5)
	list2 := List(3, 4, 5)
	vec1 := Vector(5, 4, 3)
	vec2 := Vector(3, 5, 4)
	if Reduce(intAdd, MapCar(intAdd, list1, list2)) != Reduce(intAdd, MapCar(intAdd, vec1, vec2)) {
		t.Fatalf("MapCar failed!")
	}
}

func greaterThan3(vals ...interface{}) interface{} {
	var result bool = false
	for _, val := range vals {
		v := reflect.ValueOf(val)
		intVal := v.Int()
		if intVal > 3 {
			result = true
		}
	}
	return result
}

func TestFilters(t *testing.T) {
	list1 := List(1, 2, 3, 4, 5)
	if Equal(List(1, 2, 3), RemoveIf(greaterThan3, list1)) == false {
		t.Fatalf("RemoveIf failed!")
	}
	if Equal(List(4, 5), RemoveIfNot(greaterThan3, list1)) == false {
		t.Fatalf("RemoveIfNot failed!")
	}
	if Equal(List(4, 5), Filter(greaterThan3, list1)) == false {
		t.Fatalf("Filter failed!")
	}
}

func TestReverse(t *testing.T) {
	list1 := List(3, 4, 5)
	list2 := List(5, 4, 3)
	if Equal(Reverse(list1), list2) == false {
		t.Fatalf("Map failed!")
	}
}

func TestSomeEvery(t *testing.T) {
	list1 := List(4, 5, 9, 10)
	list2 := List(4, 5, 8)
	vec1 := Vector(1, 2, 3, 4, 5, 8)
	vec2 := Vector(1, 2, 3)
	if Every(greaterThan3, list1) == false {
		t.Fatalf("Every failed!")
	}
	if Every(greaterThan3, list2) == false {
		t.Fatalf("Every failed!")
	}
	if Every(greaterThan3, vec1) != false {
		t.Fatalf("Every failed!")
	}
	if Some(greaterThan3, vec1) == false {
		t.Fatalf("Some failed!")
	}
	if Some(greaterThan3, vec2) != false {
		t.Fatalf("Some failed!")
	}
}
