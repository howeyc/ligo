// Copyright 2012 Chris Howey. All right reserved.
// Use of this source code is governed by the ISC
// license that can be found in the LICENSE file.

package ligo

// vec is the structure to simulate a clojure vector cons cell
type vec struct {
	vector []interface{}
}

func (v *vec) String() (ret string) {
	ret = "["
	ret += printSeq(v)
	ret += "]"
	return ret
}

func (v *vec) slice() []interface{} {
	return v.vector
}

// First returns the value of the first
func (v *vec) First() interface{} {
	if v != nil && len(v.vector) > 0 {
		return v.vector[0]
	}
	return nil
}

// Rest returns a new seq containing the rest
func (v *vec) Rest() Seq {
	if v != nil && len(v.vector) > 1 {
		return &vec{v.vector[1:]}
	}
	return nil
}

func (v *vec) pushNew(val interface{}) Seq {
	newvector := make([]interface{}, 0)
	newvector = append(newvector, val)
	newvector = append(newvector, v.vector...)
	return &vec{newvector}
}

func consVector(val interface{}, rest *vec) Seq {
	if rest == nil {
		newvector := make([]interface{}, 0)
		newvector = append(newvector, val)
		return &vec{newvector}
	}

	return rest.pushNew(val)
}

// Vector is a convinience function to build a vector from its arguments
func Vector(a ...interface{}) Seq {
	if len(a) == 0 {
		return nil
	}
	newvector := make([]interface{}, 0)
	newvector = append(newvector, a...)
	return &vec{newvector}
}

// MakeSeq returns a sequence from a given Go array slice
func MakeSeq(a []interface{}) Seq {
	return Vector(a...)
}
