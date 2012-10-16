// Copyright 2012 Chris Howey. All right reserved.
// Use of this source code is governed by the ISC
// license that can be found in the LICENSE file.

package ligo

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
		return cons(fn(seqvals...), MapCar(fn, newseqs...))
	}
	return cons(fn(seqvals...), nil)
}

// Every tests elements of sequences for satisfaction of a given fn. The
// first argument to fn is an element of the first sequence; each succeding
// argument is an element of a succedding sequence.
//
// Every returns false as soon as any invocation of fn returns false.
// If the end of a sequence is reached, Every returns true.
func Every(fn MultiArgFn, seqs ...Seq) bool {
	nextseqs := make([]Seq, 0)
	seqvals := make([]interface{}, 0)
	for _, seq := range seqs {
		if seq != nil {
			seqvals = append(seqvals, seq.First())
			nextseqs = append(nextseqs, seq.Rest())
		}
	}
	if len(seqvals) == 0 || len(seqvals) < len(seqs) {
		return true
	}
	return fn(seqvals...) != false && Every(fn, nextseqs...) == true
}

// Some tests elements of sequences for satisfaction of a given fn. The
// first argument to fn is an element of the first sequence; each succeding
// argument is an element of a succedding sequence.
//
// Some returns true as soon as any invocation of fn returns non-false.
// If the end of a sequence is reached, Some returns false.
func Some(fn MultiArgFn, seqs ...Seq) bool {
	nextseqs := make([]Seq, 0)
	seqvals := make([]interface{}, 0)
	for _, seq := range seqs {
		if seq != nil {
			seqvals = append(seqvals, seq.First())
			nextseqs = append(nextseqs, seq.Rest())
		}
	}
	if len(seqvals) == 0 || len(seqvals) < len(seqs) {
		return false
	}
	if fn(seqvals...) != false {
		return true
	}
	return Some(fn, nextseqs...)
}

// EqualTest returns true if the values of the sequences are equal
// based on equalityFn and the sequences are of the same length.
func EqualTest(equalityFn MultiArgFn, seqs ...Seq) bool {
	nextseqs := make([]Seq, 0)
	seqvals := make([]interface{}, 0)
	for _, seq := range seqs {
		if seq != nil {
			seqvals = append(seqvals, seq.First())
			nextseqs = append(nextseqs, seq.Rest())
		}
	}
	if len(seqvals) == 0 {
		return true
	} else if len(seqvals) < len(seqs) {
		return false
	}
	return equalityFn(seqvals...) == true && EqualTest(equalityFn, nextseqs...)
}

// RemoveIf returns a sequence from which elements that satisfy fn
// have been removed.
func RemoveIf(fn MultiArgFn, seq Seq) Seq {
	if seq == nil {
		return nil
	}
	vals := make([]interface{}, 0)
	for cell := seq; cell != nil; cell = cell.Rest() {
		val := cell.First()
		if fn(val) == false {
			vals = append(vals, val)
		}
	}
	return Vector(vals...)
}

// RemoveIfNot returns a sequence from which elements that do not
// satisfy fn have been removed.
func RemoveIfNot(fn MultiArgFn, seq Seq) Seq {
	if seq == nil {
		return nil
	}
	vals := make([]interface{}, 0)
	for cell := seq; cell != nil; cell = cell.Rest() {
		val := cell.First()
		if fn(val) != false {
			vals = append(vals, val)
		}
	}
	return Vector(vals...)
}

// Filter returns a sequence with elements that satisfy fn
func Filter(fn MultiArgFn, seq Seq) Seq {
    return RemoveIfNot(fn, seq)
}
