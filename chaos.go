package chaos

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/google/uuid"
	"golang.org/x/exp/constraints"
)

// Int returns a deterministic number between 0 and n included.
func Int[I constraints.Integer](n I, seed ...any) I {
	u := stringToSeed(seed)
	r := rand.New(rand.NewSource(u))
	return I(r.Int63n(int64(n)))
}

// Bool returns a deterministic boolean.
func Bool(seed ...any) bool {
	return stringToSeed(seed)%2 == 0
}

// Duration returns a deterministic duration between 0 and n.
func Duration(n time.Duration, seed ...any) time.Duration {
	return time.Duration(Int(n.Nanoseconds(), seed))
}

func Time(seed ...any) time.Time {
	return time.Unix(int64(Int(1<<32, seed)), 0)
}

// Float returns a deterministic float between 0 and n
func Float[F constraints.Float](n F, seed ...any) F {
	u := stringToSeed(seed)
	r := rand.New(rand.NewSource(u))
	return F(r.Float64()) * n
}

const alphanumericalChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// String returns a deterministic string of length n.
func String(length int, seed ...any) string {
	ret := ""
	for i := 0; i < length; i++ {
		index := Int(len(alphanumericalChars)-1, append(seed, i))
		ret += string(alphanumericalChars[index])
	}
	return ret
}

// SliceItem returns a deterministic item from a slice.
func SliceItem[T any](items []T, seed ...any) T {
	var t T
	if len(items) == 0 {
		return t
	}
	return items[Int(len(items)-1, seed)]
}

// UniqueSliceItems returns a slice with a length of count.
// The items are unique.
func UniqueSliceItems[T any](items []T, count int, seed ...any) ([]T, error) {
	if len(items) < count {
		return nil, fmt.Errorf("not enough items to select from: %d < %d", len(items), count)
	}

	selectedItems := make([]T, 0, count)
	availableItems := append([]T(nil), items...)
	for i := 0; i < count; i++ {
		index := Int(len(availableItems)-1, seed, i)
		selectedItems = append(selectedItems, availableItems[index])
		availableItems = append(availableItems[:index], availableItems[index+1:]...)
	}

	return selectedItems, nil
}

// MustUniqueSliceItems returns a slice with a length of count.
// The items are unique.
// If there are not enough items to select from, it panics.
func MustUniqueSliceItems[T any](items []T, count int, seed ...any) []T {
	selectedItems, err := UniqueSliceItems(items, count, seed)
	if err != nil {
		panic(err)
	}
	return selectedItems
}

// IntSlice returns a slice of deterministic numbers between 0 and high included.
// The length of the slice is length.
func IntSlice(high int, length int, seed ...any) []int {
	result := make([]int, length)
	for i := 0; i < length; i++ {
		result[i] = Int(high, seed, i)
	}

	return result
}

func UUID(seed ...any) uuid.UUID {
	s := stringToSeed(append(seed, "uuid"))
	rnd := rand.New(rand.NewSource(s))
	id, err := uuid.NewRandomFromReader(rnd)
	if err != nil {
		panic(err)
	}
	return id
}
