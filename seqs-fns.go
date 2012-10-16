// Copyright 2012 Chris Howey. All right reserved.
// Use of this source code is governed by the ISC
// license that can be found in the LICENSE file.

package ligo

// Get a copy of a Seq.
// Only the seq structure is copied, the elements of the resulting
// seq are the same as the corresponding elements of the given seq.
func CopySeq(seq Seq) Seq {
	if seq == nil {
		return nil
	} else if seq.Rest() == nil {
		return cons(seq.First(), nil)
	}
	return cons(seq.First(), CopySeq(seq.Rest()))
}

// Append returns a new seq that is the concatenation of the copies.
// Supplied seqs are left unchanged.
//      listboth := lisp.Append(list1, list2)
func Append(seqs ...Seq) Seq {
	nextseq := make([]Seq, 0)
	if len(seqs) == 0 {
		return nil
	} else if len(seqs) == 1 {
		if seqs[0] != nil {
			if seqs[0].Rest() != nil {
				nextseq = append(nextseq, seqs[0].Rest())
				return cons(seqs[0].First(), Append(nextseq...))
			}
			return cons(seqs[0].First(), nil)
		}
		return nil
	}
	if seqs[0] == nil {
		return Append(seqs[1:]...)
	}
	if seqs[0].Rest() != nil {
		nextseq = append(nextseq, Append(seqs[0].Rest()))
		nextseq = append(nextseq, seqs[1:]...)
		return cons(seqs[0].First(), Append(nextseq...))
	}
	nextseq = append(nextseq, seqs[1:]...)
	return cons(seqs[0].First(), Append(nextseq...))
}

// RevAppend constructs a copy of seq, but with elements in reverse order.
// It then appends the tail to that reversed list and returns the result.
func RevAppend(seq, tail Seq) Seq {
	if seq == nil {
		return tail
	}
	return RevAppend(seq.Rest(), cons(seq.First(), tail))
}

// Reverse constructs a copy of seq, but with elements in reverse order.
func Reverse(seq Seq) Seq {
	return RevAppend(seq, nil)
}

// Subseq creates a copy of seq bounded by start and end
func SubSeq(seq Seq, start, end uint) Seq {
	if seq == nil || start >= end {
		return nil
	}
	if seq.Rest() == nil {
		return seq
	}
	if start > 0 {
		return SubSeq(seq.Rest(), start-1, end-1)
	} else if start == 0 {
		return cons(seq.First(), SubSeq(seq.Rest(), start, end-1))
	}
	return nil
}

// Takes forst n elements of a seq
func Take(seq Seq, n uint) Seq {
	return SubSeq(seq, 0, n)
}

// Drop forst n elements of a seq
func Drop(seq Seq, n uint) Seq {
	if n == 0 || seq == nil {
		return seq
	}
	return Drop(seq.Rest(), n-1)
}

// Get the length of a given seq
func Length(seq Seq) int {
	return len(GetSlice(seq))
}
