package chaos

import (
	"fmt"
	"math/rand"
	"time"

	"golang.org/x/exp/constraints"
)

// Int returns a deterministic number between 0 and n included.
func Int[I constraints.Integer](seed string, n I) I {
	u := stringToSeed(seed)
	r := rand.New(rand.NewSource(u))
	return I(r.Int63n(int64(n)))
}

// Bool returns a deterministic boolean.
func Bool(seed string) bool {
	return stringToSeed(seed)%2 == 0
}

// Duration returns a deterministic duration between 0 and n.
func Duration(seed string, n time.Duration) time.Duration {
	return time.Duration(Int(seed, n.Nanoseconds()))
}

// Float returns a deterministic float between 0 and n
func Float[F constraints.Float](seed string, n F) F {
	u := stringToSeed(seed)
	r := rand.New(rand.NewSource(u))
	return F(r.Float64()) * n
}

const alphanumericalChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// String returns a deterministic string of length n.
func String(seed string, length int) string {
	ret := ""
	for i := 0; i < length; i++ {
		iSeed := fmt.Sprintf("%s-%d", seed, i)
		index := Int(iSeed, len(alphanumericalChars)-1)
		ret += string(alphanumericalChars[index])
	}
	return ret
}

// SliceItem returns a deterministic item from a slice.
func SliceItem[T any](seed string, items []T) T {
	var t T
	if len(items) == 0 {
		return t
	}
	return items[Int(seed, len(items)-1)]
}

// UniqueSliceItems returns a slice with a length of count.
// The items are unique.
func UniqueSliceItems[T any](seed string, items []T, count int) ([]T, error) {
	if len(items) < count {
		return nil, fmt.Errorf("not enough items to select from: %d < %d", len(items), count)
	}

	selectedItems := make([]T, 0, count)
	availableItems := append([]T(nil), items...)
	for i := 0; i < count; i++ {
		index := Int(fmt.Sprintf("%s-%d", seed, i), len(availableItems)-1)
		selectedItems = append(selectedItems, availableItems[index])
		availableItems = append(availableItems[:index], availableItems[index+1:]...)
	}

	return selectedItems, nil
}

// MustUniqueSliceItems returns a slice with a length of count.
// The items are unique.
// If there are not enough items to select from, it panics.
func MustUniqueSliceItems[T any](seed string, items []T, count int) []T {
	selectedItems, err := UniqueSliceItems(seed, items, count)
	if err != nil {
		panic(err)
	}
	return selectedItems
}

// IntSlice returns a slice of deterministic numbers between 0 and high included.
// The length of the slice is length.
func IntSlice(seed string, high int, length int) []int {
	result := make([]int, length)
	for i := 0; i < length; i++ {
		result[i] = Int(fmt.Sprintf("%s-%d", seed, i), high)
	}

	return result
}
