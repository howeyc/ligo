// Copyright 2012 Chris Howey. All right reserved.
// Use of this source code is governed by the ISC
// license that can be found in the LICENSE file.

package ligo

// pair is the structure to simulate a lisp cons cell
type pair struct {
	car interface{}
	cdr Seq
}

// First returns the value of the pair
func (p *pair) First() interface{} {
	if p != nil {
		return p.car
	}
	return nil
}

// Rest returns the cdr of the pair
func (p *pair) Rest() Seq {
	if p != nil {
		return p.cdr
	}
	return nil
}

func consPair(val interface{}, rest *pair) Seq {
	return &pair{val, rest}
}

// List is a convinience function to build a list from its arguments
func List(a ...interface{}) Seq {
	if len(a) == 0 {
		return nil
	} else if len(a) == 1 {
		return consPair(a[0], nil)
	}
	start := cons(a[len(a)-1], nil)
	for i := len(a) - 2; i >= 0; i-- {
		start = cons(a[i], start)
	}
	return start
}
