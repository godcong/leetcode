package _0284

import (
	. "github.com/godcong/leetcode/common"
)

/*   Below is the interface for Iterator, which is already defined for you.
 *
 *   type Iterator struct {
 *
 *   }
 *
 *   func (this *Iterator) hasNext() bool {
 *		// Returns true if the iteration has more elements.
 *   }
 *
 *   func (this *Iterator) next() int {
 *		// Returns the next element in the iteration.
 *   }
 */

type PeekingIterator struct {
	iter     *Iterator
	_hasNext bool
	_next    int
}

func Constructor(iter *Iterator) *PeekingIterator {
	return &PeekingIterator{iter: iter, _hasNext: iter.hasNext(), _next: iter.next()}
}

func (this *PeekingIterator) hasNext() bool {
	return this._hasNext
}

func (this *PeekingIterator) next() int {
	ret := this._next
	this._hasNext = this.iter.hasNext()
	if this._hasNext {
		this._next = this.iter.next()
	}
	return ret
}

func (this *PeekingIterator) peek() int {
	return this._next
}
