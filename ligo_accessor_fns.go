// Copyright 2012 Chris Howey. All right reserved.
// Use of this source code is governed by the ISC
// license that can be found in the LICENSE file.

package ligo

// NthRest returns the sequence at the "index" n. n <= 0 returns the passed
// in sequence.
func NthRest(seq Seq, n uint) Seq {
    return Drop(seq, n)
}

// Nth returns the value at the given "index" n of the sequence seq.
// n of zero is equivalent to First.
func Nth(seq Seq, n uint) interface {} {
    return First(NthRest(seq, n))
}

// First returns the value at the beginning of the sequence.
func First(seq Seq) interface{} {
    if seq == nil {
        return nil
    }
    return seq.First()
}

// Second returns the second value of the sequence.
func Second(seq Seq) interface{} {
    return Nth(seq, 1)
}

// Third returns the third value of the sequence.
func Third(seq Seq) interface{} {
    return Nth(seq, 2)
}

// Fourth returns the fourth value of the sequence.
func Fourth(seq Seq) interface{} {
    return Nth(seq, 3)
}

// Fifth returns the fifth value of the sequence.
func Fifth(seq Seq) interface{} {
    return Nth(seq, 4)
}
