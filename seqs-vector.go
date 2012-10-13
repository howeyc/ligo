// Copyright 2012 Chris Howey. All right reserved.
// Use of this source code is governed by the ISC
// license that can be found in the LICENSE file.

package ligo

// vec is the structure to simulate a clojure vector cons cell
type vec struct {
	vector []interface{}
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
	} else if len(a) == 1 {
		return consVector(a[0], nil)
	}
	start := cons(a[len(a)-1], nil)
	for i := len(a) - 2; i >= 0; i-- {
		start = cons(a[i], start)
	}
	return start
}
