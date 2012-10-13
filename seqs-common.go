// Copyright 2012 Chris Howey. All right reserved.
// Use of this source code is governed by the ISC
// license that can be found in the LICENSE file.

// Package ligo provides the basics of Lisp In Go.
package ligo

import (
	"reflect"
	"unsafe"
)

// Seq is the interface the specifies the basic functions of a sequence
//
// First returns the value at the start of the sequence
// Rest returns another sequence containing all elements after the first
type Seq interface {
	First() interface{}
	Rest() Seq
}

// Returns a new cons cell
// Cons can be used to build a list:
//      list1 = Cons(1, Cons(2, Cons(3, nil)))
//      list2 = Cons(5, list2)
func cons(val interface{}, rest Seq) Seq {
	// Default is vector
	if rest == nil {
		return consVector(val, nil)
	}
	vrest := reflect.ValueOf(rest).Elem()

	// Pair??
	if vrest.Type().String() == "ligo.pair" {
		return &pair{val, rest}
	}

	// Vec??
	if vrest.Type().String() == "ligo.vec" {
		restaddr := vrest.UnsafeAddr()
		restptr := unsafe.Pointer(restaddr)
		return consVector(val, (*vec)(restptr))
	}

	return nil
}

// GetSlice returns the given seq as a Go array slice
func GetSlice(seq Seq) []interface{} {
	if seq == nil {
		return nil
	}
	vseq := reflect.ValueOf(seq).Elem()

	// Vec??
	if vseq.Type().String() == "ligo.vec" {
		vecaddr := vseq.UnsafeAddr()
		vecptr := unsafe.Pointer(vecaddr)
		return (*vec)(vecptr).slice()
	}

	vector := make([]interface{}, 0)
	cell := seq
	for cell != nil {
		vector = append(vector, cell.First())
		cell = cell.Rest()
	}
	return vector
}
