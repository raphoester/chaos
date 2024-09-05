package chaos

import (
	"errors"
	"fmt"
)

// NewSliceProcessor returns a new SliceProcessor.
func NewSliceProcessor[S ~[]T, T any](c *Chaos) *SliceProcessor[S, T] {
	return &SliceProcessor[S, T]{
		c: c,
	}
}

// SliceProcessor is a helper to extract values from slices.
type SliceProcessor[S ~[]T, T any] struct {
	c *Chaos
}

func SliceItem[S ~[]T, T any](items S) T {
	return NewSliceProcessor[S, T](singleton).Item(items)
}

// Item returns a random item from the slice.
func (s *SliceProcessor[S, T]) Item(items S) T {
	var ret T
	if len(items)== 0 {
		return ret
	}

	return items[s.c.Int(len(items)-1)]
}

var (
	ErrNotEnoughItemsInSlice = errors.New("slice does not contain enough items")
)

// UniqueSliceItems returns a random item from the slice.
func UniqueSliceItems[S ~[]T, T any](items S, count int) (S, error) {
	return NewSliceProcessor[S, T](singleton).UniqueItems(items, count)
}

// UniqueItems returns a slice with a length of count.
// The items are unique.
// If there are not enough items to select from, it returns an error.
func (s *SliceProcessor[S, T]) UniqueItems(items S, count int) (S, error) {
	if len(items) < count {
		return nil, errors.Join(ErrNotEnoughItemsInSlice,
			fmt.Errorf("not enough items to select from: %d < %d", len(items), count))
	}

	selectedItems := make(S, 0, count)
	availableItems := append(S(nil), items...)
	for i := 0; i < count; i++ {
		index := s.c.Int(len(availableItems) - 1)
		selectedItems = append(selectedItems, availableItems[index])
		availableItems = append(availableItems[:index], availableItems[index+1:]...)
	}

	return selectedItems, nil
}

// MustUniqueSliceItems returns a random item from the slice.
func MustUniqueSliceItems[S ~[]T, T any](items S, count int) S {
	return NewSliceProcessor[S, T](singleton).MustUniqueItems(items, count)
}

// MustUniqueItems returns a slice with a length of count.
// The items are unique.
// If there are not enough items to select from, it panics.
func (s *SliceProcessor[S, T]) MustUniqueItems(items S, count int) S {
	selectedItems, err := s.UniqueItems(items, count)
	if err != nil {
		panic(err)
	}
	return selectedItems
}
