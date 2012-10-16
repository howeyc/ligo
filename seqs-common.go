// Copyright 2012 Chris Howey. All right reserved.
// Use of this source code is governed by the ISC
// license that can be found in the LICENSE file.

// Package ligo provides the basics of Lisp In Go.
package ligo

import (
	"fmt"
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
	restaddr := vrest.UnsafeAddr()
	restptr := unsafe.Pointer(restaddr)

	// Pair??
	if vrest.Type().String() == "ligo.pair" {
		return consPair(val, (*pair)(restptr))
	}

	// Vec??
	if vrest.Type().String() == "ligo.vec" {
		return consVector(val, (*vec)(restptr))
	}

	return nil
}

// ToSlice returns the given seq as a Go array slice
func ToSlice(seq Seq) []interface{} {
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

func printSeq(seq Seq) (ret string) {
	if seq == nil {
		return ""
	}

	ret = fmt.Sprintf("%v", seq.First())
	cell := seq.Rest()
	for cell != nil {
		ret += fmt.Sprintf(" %v", cell.First())
		cell = cell.Rest()
	}
	return ret
}

func valEquality(vals ...interface{}) interface{} {
	var result bool = true
	prevVal := vals[0]
	for _, val := range vals {
		if reflect.DeepEqual(prevVal, val) != true {
			result = false
			break
		}
		prevVal = val
	}
	return result
}

// Equal returns true if the values of the sequences are equal
// (reflect.DeepEqual) and the sequences are of the same length.
func Equal(seqs ...Seq) bool {
	return EqualTest(valEquality, seqs...)
}
