package iterator

import (
	"errors"
	"fmt"

	"github.com/Xuanwo/storage/types"
)

// ErrDone means this iterator has been done.
var ErrDone = errors.New("iterator is done")

// Iterator is the iterator interface used to do iterate
type Iterator interface {
	Next() (types.Informer, error)
}

// NextFunc is the func used in iterator.
type NextFunc func(*[]types.Informer) error

// PrefixBasedIterator is the prefix based iterator.
type PrefixBasedIterator struct {
	buf   []types.Informer
	index int
	next  NextFunc
}

// NewPrefixBasedIterator will return a new prefix based iterator.
func NewPrefixBasedIterator(fn NextFunc) *PrefixBasedIterator {
	return &PrefixBasedIterator{
		buf:   nil,
		index: 0,
		next:  fn,
	}
}

// Next implements Iterator interface.
//
// Next call is not thread safe, do not call it in multi goroutine.
func (it *PrefixBasedIterator) Next() (i types.Informer, err error) {
	if it.index < len(it.buf) {
		it.index++
		return it.buf[it.index-1], nil
	}

	err = it.next(&it.buf)
	if err == nil {
		it.index = 1
		return it.buf[0], nil
	}
	if errors.Is(err, ErrDone) {
		return nil, ErrDone
	}
	return nil, fmt.Errorf("iterator next failed: %w", err)
}
