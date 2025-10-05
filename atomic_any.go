package atomicany

import (
	"sync/atomic"
)

type AtomicAny[T any] struct{ v atomic.Value }

func NewAtomicAny[T any](v T) *AtomicAny[T] {
	aa := &AtomicAny[T]{}
	aa.v.Store(v)
	return aa
}

func (aa *AtomicAny[T]) Load() (v T) {
	a := aa.v.Load()
	if a != nil {
		return a.(T)
	}
	return
}

func (aa *AtomicAny[T]) Store(v T) { aa.v.Store(v) }
