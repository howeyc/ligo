// Copyright 2012 Chris Howey. All right reserved.
// Use of this source code is governed by the ISC
// license that can be found in the LICENSE file.

package ligo

// pair is the structure to simulate a lisp cons cell
type pair struct {
	car interface{}
	cdr Seq
}

func (p *pair) String() (ret string) {
	ret = "("
	ret += printSeq(p)
	ret += ")"
	return ret
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
	// From what I can see it looks that if you pass nil as rest to 
	// consPair, Go will pass a pointer to an empty pair struct in its
	// place. So in some cases instead of cdr being nil, it points to a pair
	// with a nil car.
	if p != nil && p.cdr != nil && p.cdr.First() != nil {
		return p.cdr
	}
	return nil
}

func consPair(val interface{}, rest *pair) Seq {
	return &pair{val, rest}
}

// List is a function to build a list sequence from its arguments
func List(a ...interface{}) Seq {
	if len(a) == 0 {
		return nil
	}
	start := consPair(a[len(a)-1], nil)
	for i := len(a) - 2; i >= 0; i-- {
		start = cons(a[i], start)
	}
	return start
}
