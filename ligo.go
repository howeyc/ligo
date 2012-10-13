// Copyright 2012 Chris Howey. All right reserved.
// Use of this source code is governed by the ISC
// license that can be found in the LICENSE file.

// Package ligo provides the basics of Lisp In Go.
package ligo

// Seq is the interface the specifies the basic functions of a sequence
//
// First returns the value at the start of the sequence
// Rest returns another sequence containing all elements after the first
type Seq interface {
	First() interface{}
	Rest() Seq
}

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

// The fundamental Cons function
//
// Returns a fresh pair of which val is the car
// and rest is the cdr
//
// Cons can be used to build a list:
//      list1 = Cons(1, Cons(2, Cons(3, nil)))
//      list2 = Cons(5, list2)
func Cons(val interface{}, rest Seq) *pair {
	return &pair{val, rest}
}

// List is a convinience function to build a list from its arguments
//
// The following lists are equivalent:
//      list1 = Cons(1, Cons(2, Cons(3, nil)))
//      list2 = List(1, 2, 3)
func List(a ...interface{}) Seq {
	if len(a) == 0 {
		return nil
	} else if len(a) == 1 {
		return Cons(a[0], nil)
	}
	start := Cons(a[len(a)-1], nil)
	for i := len(a) - 2; i >= 0; i-- {
		start = Cons(a[i], start)
	}
	return start
}

// All functions passed to high-order functions in the package
// must be able to receive zero (one-arg as nil), one, or many
// arguments with a single return value.
// The functions must use reflect package to retrieve values.
type MultiArgFn func(...interface{}) interface{}

// Reduce is a uses a function fn to combine the elements of seq
//      Reduce(AddInts, List(1, 2, 4, 8)) => 15
//      Reduce(MultiInts, List(2, 2, 4)) => 16
// Note that the return type will be the whatever is returned by
// the provided function fn.
func Reduce(fn MultiArgFn, seq Seq) interface{} {
	if seq != nil {
		if seq.Rest() != nil {
			return fn(seq.First(), Reduce(fn, seq.Rest()))
		}
		return fn(seq.First())
	}
	return nil
}

// MapCar involves applying fn to successive sets of arguments in
// which one argument is obtained from each seq.
// fn is applied to the first element of each seq, then to the second
// element of each seq, and so on. The iteration terminates when the
// shortest list runs out, and excess elements in other lists are ignored
// The results of each successive application of fn are returned in a
// new Seq.
//      MapCar(AddTen, List(1, 2, 3, 4)) => (11, 12, 13, 14)
func MapCar(fn MultiArgFn, seqs ...Seq) Seq {
	newseqs := make([]Seq, 0)
	seqvals := make([]interface{}, 0)
	for _, seq := range seqs {
        if seq == nil {
            return nil
        }
		seqvals = append(seqvals, seq.First())
		newseqs = append(newseqs, seq.Rest())
	}
	if len(newseqs) > 0 {
		return Cons(fn(seqvals...), MapCar(fn, newseqs...))
	}
	return Cons(fn(seqvals...), nil)
}

// Get a copy of a Seq.
// Only the seq structure is copied, the elements of the resulting
// seq are the same as the corresponding elements of the given seq.
func CopySeq(seq Seq) Seq {
	if seq == nil {
		return nil
	} else if seq.Rest() == nil {
		return Cons(seq.First(), nil)
	}
	return Cons(seq.First(), CopySeq(seq.Rest()))
}

// Append returns a new seq that is the concatenation of the copies.
// Supplied seqs are left unchanged.
//      listboth := lisp.Append(list1, list2)
func Append(seqs ...Seq) Seq {
    nextseq := make([]Seq, 0)
	if len(seqs) == 0 {
		return nil
	} else if len(seqs) == 1 {
		if seqs[0].Rest() != nil {
			nextseq = append(nextseq, seqs[0].Rest())
			nextseq = append(nextseq, seqs[1:]...)
			return Cons(seqs[0].First(), Append(nextseq...))
		}
		return Cons(seqs[0].First(), nil)
	}
    if seqs[0].Rest() != nil {
	    nextseq = append(nextseq, Append(seqs[0].Rest()))
	    nextseq = append(nextseq, seqs[1:]...)
        return Cons(seqs[0].First(), Append(nextseq...))
    }
	nextseq = append(nextseq, seqs[1:]...)
	return Cons(seqs[0].First(), Append(nextseq...))
}
