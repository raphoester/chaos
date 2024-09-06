package chaos_test

import (
	"fmt"
	"testing"

	"github.com/raphoester/chaos"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSliceItem(t *testing.T) {
	t.Run("deterministic output", func(t *testing.T) {
		c := chaos.New(t.Name())
		c.Fix()
		p := chaos.NewSliceProcessor[[]int](c)

		items := []int{1, 2, 3, 4, 5}
		assert.Equal(t, p.Item(items), p.Item(items))
	})

	t.Run("selects all items with equal probability", func(t *testing.T) {
		c := chaos.New(t.Name())
		p := chaos.NewSliceProcessor[[]int](c)
		items := []int{1, 2, 3, 4, 5}
		results := make(map[int]int)
		const iterations = 10000
		for i := 0; i < iterations; i++ {
			results[p.Item(items)]++
		}
		assert.Len(t, results, len(items), "Expected to select all items")
		expectedCount := iterations / len(items)
		tolerance := float64(expectedCount) * 0.1 // 10% tolerance
		for _, count := range results {
			assert.InDelta(t, expectedCount, count, tolerance, "Expected roughly equal distribution")
		}
	})

	t.Run("returns zero value for empty slice", func(t *testing.T) {
		c := chaos.New(t.Name())
		pInt := chaos.NewSliceProcessor[[]int](c)
		pString := chaos.NewSliceProcessor[[]string, string](c)

		var emptyIntSlice []int
		assert.Equal(t, 0, pInt.Item(emptyIntSlice))

		var emptyStringSlice []string
		assert.Equal(t, "", pString.Item(emptyStringSlice))
	})

	t.Run("selects from all items", func(t *testing.T) {
		c := chaos.New(t.Name())
		p := chaos.NewSliceProcessor[[]int](c)

		items := []int{1, 2, 3, 4, 5}
		selected := make(map[int]bool)
		for i := 0; i < 1000; i++ {
			selected[p.Item(items)] = true
		}
		assert.Len(t, selected, len(items))
	})

	t.Run("edge case: single item slice", func(t *testing.T) {
		c := chaos.New(t.Name())
		p := chaos.NewSliceProcessor[[]int](c)
		items := []int{42}
		assert.Equal(t, 42, p.Item(items))
	})

	t.Run("selects all items with approximately equal probability", func(t *testing.T) {
		c := chaos.New(t.Name())
		p := chaos.NewSliceProcessor[[]int](c)
		items := []int{1, 2, 3, 4, 5}
		results := make(map[int]int)
		const iterations = 100000
		for i := 0; i < iterations; i++ {
			results[p.Item(items)]++
		}
		assert.Len(t, results, len(items), "Expected to select all items")

		expectedCount := iterations / len(items)
		tolerance := 0.05 // 5% tolerance
		for item, count := range results {
			lowerBound := float64(expectedCount) * (1 - tolerance)
			upperBound := float64(expectedCount) * (1 + tolerance)
			assert.GreaterOrEqual(t, float64(count), lowerBound, "Count for item %d is lower than expected", item)
			assert.LessOrEqual(t, float64(count), upperBound, "Count for item %d is higher than expected", item)
		}
	})
}

func TestUniqueSliceItems(t *testing.T) {
	t.Run("deterministic output", func(t *testing.T) {
		c := chaos.New(t.Name())
		c.Fix()
		p := chaos.NewSliceProcessor[[]int](c)

		items := []int{1, 2, 3, 4, 5}
		result1, err1 := p.UniqueItems(items, 3)
		result2, err2 := p.UniqueItems(items, 3)

		require.NoError(t, err1)
		require.NoError(t, err2)
		assert.Equal(t, result1, result2)
	})

	t.Run("returns error when not enough items", func(t *testing.T) {
		c := chaos.New(t.Name())
		c.Fix()
		p := chaos.NewSliceProcessor[[]int](c)

		items := []int{1, 2, 3}
		_, err := p.UniqueItems(items, 4)
		assert.ErrorIs(t, err, chaos.ErrNotEnoughItemsInSlice)
	})

	t.Run("returns unique items", func(t *testing.T) {
		c := chaos.New(t.Name())
		p := chaos.NewSliceProcessor[[]int](c)

		items := []int{1, 2, 3, 4, 5}
		result, err := p.UniqueItems(items, 3)
		require.NoError(t, err)
		assert.Len(t, result, 3)
		assert.Len(t, uniqueInts(result), 3)
	})

	t.Run("different seeds produce unique results up to maximum possible combinations", func(t *testing.T) {
		c := chaos.New(t.Name())
		p := chaos.NewSliceProcessor[[]int](c)

		items := []int{1, 2, 3, 4, 5}
		results := make(map[string]bool)
		maxCombinations := 60
		iterations := 1000

		for i := 0; i < iterations; i++ {
			result, err := p.UniqueItems(items, 3)
			require.NoError(t, err)
			key := fmt.Sprintf("%v", result)
			results[key] = true
		}

		assert.Len(t, results, maxCombinations, "Expected all possible unique combinations")

		c.Fix()
		// Check if the function is deterministic for the same seed
		result1, _ := p.UniqueItems(items, 3)
		result2, _ := p.UniqueItems(items, 3)
		assert.Equal(t, result1, result2, "Expected same result for the same seed")

		c.Unfix()
		// Check if different seeds produce different results
		result3, _ := p.UniqueItems(items, 3)
		assert.NotEqual(t, result1, result3, "Expected different results for different seeds")
	})

	t.Run("edge case: request all items", func(t *testing.T) {
		c := chaos.New(t.Name())
		p := chaos.NewSliceProcessor[[]int](c)

		items := []int{1, 2, 3, 4, 5}
		result, err := p.UniqueItems(items, 5)
		require.NoError(t, err)
		assert.ElementsMatch(t, items, result)
	})
}

// Helper function to get unique integers from a slice
func uniqueInts(slice []int) map[int]struct{} {
	unique := make(map[int]struct{})
	for _, num := range slice {
		unique[num] = struct{}{}
	}
	return unique
}
