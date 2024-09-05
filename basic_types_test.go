package chaos_test

import (
	"fmt"
	"testing"
	"time"

	"math"

	"github.com/google/uuid"
	"github.com/raphoester/chaos"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInt(t *testing.T) {

	t.Run("fixed chaos produces the same output", func(t *testing.T) {
		c := chaos.New(t.Name())
		c.Fix()
		assert.Equal(t, c.Int(100), c.Int(100))
	})

	t.Run("unfixed chaos produces unique results", func(t *testing.T) {
		c := chaos.New(t.Name())
		results := make(map[int]bool)
		for i := 0; i < 1000; i++ {
			result := c.Int(math.MaxInt32)
			assert.False(t, results[result], "Expected unique result for each seed")
			results[result] = true
		}
		assert.Len(t, results, 1000, "Expected 1000 unique results")
	})

	t.Run("respects upper bound", func(t *testing.T) {
		c := chaos.New(t.Name())
		for i := 0; i < 1000; i++ {
			result := c.Int(10)
			assert.GreaterOrEqual(t, result, 0)
			assert.LessOrEqual(t, result, 10)
		}
	})

	t.Run("edge case: zero", func(t *testing.T) {
		c := chaos.New(t.Name())
		assert.Equal(t, 0, c.Int(0))
	})

	t.Run("edge case: negative upper bound", func(t *testing.T) {
		c := chaos.New(t.Name())
		result := c.Int(-10)
		assert.GreaterOrEqual(t, result, -10)
		assert.LessOrEqual(t, result, 0)
	})

	t.Run("edge case: MaxInt32", func(t *testing.T) {
		c := chaos.New(t.Name())
		result := c.Int(math.MaxInt32)
		assert.GreaterOrEqual(t, result, 0)
		assert.LessOrEqual(t, result, math.MaxInt32)
	})
}

func TestBool(t *testing.T) {
	t.Run("deterministic output", func(t *testing.T) {
		c := chaos.New(t.Name())
		c.Fix()
		assert.Equal(t, c.Bool(), c.Bool())
	})

	t.Run("different seeds produce both true and false", func(t *testing.T) {
		c := chaos.New(t.Name())
		results := make(map[bool]int)
		for i := 0; i < 1000; i++ {
			result := c.Bool()
			results[result]++
		}
		assert.Len(t, results, 2, "Expected both true and false results")
		assert.Greater(t, results[true], 0, "Expected some true results")
		assert.Greater(t, results[false], 0, "Expected some false results")
	})
}

func TestDuration(t *testing.T) {
	t.Run("deterministic output", func(t *testing.T) {
		c := chaos.New(t.Name())
		c.Fix()
		assert.Equal(t, c.Duration(time.Hour), c.Duration(time.Hour))
	})

	t.Run("different seeds produce unique results", func(t *testing.T) {
		c := chaos.New(t.Name())
		results := make(map[time.Duration]bool)
		for i := 0; i < 1000; i++ {
			result := c.Duration(time.Hour)
			assert.False(t, results[result], "Expected unique result for each seed")
			results[result] = true
		}
		assert.Len(t, results, 1000, "Expected 1000 unique results")
	})

	t.Run("respects upper bound", func(t *testing.T) {
		c := chaos.New(t.Name())
		for i := 0; i < 1000; i++ {
			result := c.Duration(time.Hour)
			assert.GreaterOrEqual(t, result, time.Duration(0))
			assert.LessOrEqual(t, result, time.Hour)
		}
	})

	t.Run("edge case: zero duration", func(t *testing.T) {
		c := chaos.New(t.Name())
		assert.Equal(t, time.Duration(0), c.Duration(time.Duration(0)))
	})

	t.Run("edge case: negative duration", func(t *testing.T) {
		c := chaos.New(t.Name())
		result := c.Duration(-time.Hour)
		assert.GreaterOrEqual(t, result, -time.Hour)
		assert.LessOrEqual(t, result, time.Duration(0))
	})
}

func TestTime(t *testing.T) {
	t.Run("deterministic output", func(t *testing.T) {
		c := chaos.New(t.Name())
		c.Fix()
		assert.Equal(t, c.Time(), c.Time())
	})

	t.Run("different seeds produce unique results", func(t *testing.T) {
		c := chaos.New(t.Name())
		results := make(map[time.Time]bool)
		for i := 0; i < 1000; i++ {
			result := c.Time()
			assert.False(t, results[result], "Expected unique result for each seed")
			results[result] = true
		}
		assert.Len(t, results, 1000, "Expected 1000 unique results")
	})

	t.Run("respects time range", func(t *testing.T) {
		c := chaos.New(t.Name())
		minTime := time.Unix(0, 0)
		maxTime := time.Unix(1<<32-1, 0)
		for i := 0; i < 1000; i++ {
			result := c.Time()
			assert.True(t, result.After(minTime) || result.Equal(minTime))
			assert.True(t, result.Before(maxTime) || result.Equal(maxTime))
		}
	})

	t.Run("edge case: far future time", func(t *testing.T) {
		c := chaos.New(t.Name())
		farFuture := time.Now().AddDate(100, 0, 0)
		result := c.Time()
		assert.True(t, result.Before(farFuture) || result.Equal(farFuture))
	})
}

func TestFloat64(t *testing.T) {
	t.Run("deterministic output", func(t *testing.T) {
		c := chaos.New(t.Name())
		c.Fix()
		assert.Equal(t, c.Float64(1.0), c.Float64(1.0))
	})

	t.Run("different seeds produce unique results", func(t *testing.T) {
		c := chaos.New(t.Name())
		results := make(map[float64]bool)
		for i := 0; i < 1000; i++ {
			result := c.Float64(1.0)
			assert.False(t, results[result], "Expected unique result for each seed")
			results[result] = true
		}
		assert.Len(t, results, 1000, "Expected 1000 unique results")
	})

	t.Run("respects upper bound", func(t *testing.T) {
		c := chaos.New(t.Name())
		for i := 0; i < 1000; i++ {
			result := c.Float64(10.0)
			assert.GreaterOrEqual(t, result, 0.0)
			assert.Less(t, result, 10.0)
		}
	})

	t.Run("edge case: zero", func(t *testing.T) {
		c := chaos.New(t.Name())
		assert.Equal(t, 0.0, c.Float64(0.0))
	})

	t.Run("edge case: negative upper bound", func(t *testing.T) {
		c := chaos.New(t.Name())
		result := c.Float64(-10.0)
		assert.GreaterOrEqual(t, result, -10.0)
		assert.Less(t, result, 0.0)
	})

	t.Run("edge case: very small positive number", func(t *testing.T) {
		c := chaos.New(t.Name())
		result := c.Float64(1e-10)
		assert.GreaterOrEqual(t, result, 0.0)
		assert.Less(t, result, 1e-10)
	})
}

func TestString(t *testing.T) {
	t.Run("deterministic output", func(t *testing.T) {
		c := chaos.New(t.Name())
		c.Fix()
		assert.Equal(t, c.String(10), c.String(10))
	})

	t.Run("different seeds produce unique results", func(t *testing.T) {
		c := chaos.New(t.Name())
		results := make(map[string]bool)
		for i := 0; i < 1000; i++ {
			result := c.String(10)
			assert.False(t, results[result], "Expected unique result for each seed")
			results[result] = true
		}
		assert.Len(t, results, 1000, "Expected 1000 unique results")
	})

	t.Run("respects length", func(t *testing.T) {
		c := chaos.New(t.Name())
		for i := 0; i < 100; i++ {
			result := c.String(i)
			assert.Len(t, result, i)
		}
	})

	t.Run("contains only alphanumeric characters", func(t *testing.T) {
		c := chaos.New(t.Name())
		result := c.String(1000)
		for _, char := range result {
			assert.NotContains(t, "!@#$%^&*()¡™£¢∞§¶•ªº", string(char))
		}
	})

	t.Run("edge case: empty string", func(t *testing.T) {
		c := chaos.New(t.Name())
		assert.Empty(t, c.String(0))
	})
}

func TestSliceItem(t *testing.T) {
	t.Run("deterministic output", func(t *testing.T) {
		c := chaos.New(t.Name())
		c.Fix()
		p := chaos.NewSliceProcessor[[]int, int](c)

		items := []int{1, 2, 3, 4, 5}
		assert.Equal(t, p.Item(items), p.Item(items))
	})

	t.Run("selects all items with equal probability", func(t *testing.T) {
		c := chaos.New(t.Name())
		p := chaos.NewSliceProcessor[[]int, int](c)
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
		pInt := chaos.NewSliceProcessor[[]int, int](c)
		pString := chaos.NewSliceProcessor[[]string, string](c)

		var emptyIntSlice []int
		assert.Equal(t, 0, pInt.Item(emptyIntSlice))

		var emptyStringSlice []string
		assert.Equal(t, "", pString.Item(emptyStringSlice))
	})

	t.Run("selects from all items", func(t *testing.T) {
		c := chaos.New(t.Name())
		p := chaos.NewSliceProcessor[[]int, int](c)

		items := []int{1, 2, 3, 4, 5}
		selected := make(map[int]bool)
		for i := 0; i < 1000; i++ {
			selected[p.Item(items)] = true
		}
		assert.Len(t, selected, len(items))
	})

	t.Run("edge case: single item slice", func(t *testing.T) {
		c := chaos.New(t.Name())
		p := chaos.NewSliceProcessor[[]int, int](c)
		items := []int{42}
		assert.Equal(t, 42, p.Item(items))
	})

	t.Run("selects all items with approximately equal probability", func(t *testing.T) {
		c := chaos.New(t.Name())
		p := chaos.NewSliceProcessor[[]int, int](c)
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
		p := chaos.NewSliceProcessor[[]int, int](c)

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
		p := chaos.NewSliceProcessor[[]int, int](c)

		items := []int{1, 2, 3}
		_, err := p.UniqueItems(items, 4)
		assert.ErrorIs(t, err, chaos.ErrNotEnoughItemsInSlice)
	})

	t.Run("returns unique items", func(t *testing.T) {
		c := chaos.New(t.Name())
		p := chaos.NewSliceProcessor[[]int, int](c)

		items := []int{1, 2, 3, 4, 5}
		result, err := p.UniqueItems(items, 3)
		require.NoError(t, err)
		assert.Len(t, result, 3)
		assert.Len(t, uniqueInts(result), 3)
	})

	t.Run("different seeds produce unique results up to maximum possible combinations", func(t *testing.T) {
		c := chaos.New(t.Name())
		p := chaos.NewSliceProcessor[[]int, int](c)

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
		p := chaos.NewSliceProcessor[[]int, int](c)

		items := []int{1, 2, 3, 4, 5}
		result, err := p.UniqueItems(items, 5)
		require.NoError(t, err)
		assert.ElementsMatch(t, items, result)
	})
}

func TestIntSlice(t *testing.T) {
	t.Run("deterministic output", func(t *testing.T) {
		c := chaos.New(t.Name())
		c.Fix()
		assert.Equal(t, c.IntSlice(10, 5), c.IntSlice(10, 5))
	})

	t.Run("different seeds produce unique results", func(t *testing.T) {
		c := chaos.New(t.Name())
		results := make(map[string]bool)
		for i := 0; i < 1000; i++ {
			result := c.IntSlice(100, 5)
			key := fmt.Sprintf("%v", result)
			assert.False(t, results[key], "Expected unique result for each seed")
			results[key] = true
		}
		assert.Len(t, results, 1000, "Expected 1000 unique results")
	})

	t.Run("respects length", func(t *testing.T) {
		c := chaos.New(t.Name())
		result := c.IntSlice(10, 5)
		assert.Len(t, result, 5)
	})

	t.Run("respects upper bound", func(t *testing.T) {
		c := chaos.New(t.Name())
		result := c.IntSlice(10, 1000)
		for _, num := range result {
			assert.GreaterOrEqual(t, num, 0)
			assert.LessOrEqual(t, num, 10)
		}
	})

	t.Run("edge case: zero length", func(t *testing.T) {
		c := chaos.New(t.Name())
		assert.Empty(t, c.IntSlice(10, 0))
	})
}

func TestUUID(t *testing.T) {
	t.Run("deterministic output", func(t *testing.T) {
		c := chaos.New(t.Name())
		c.Fix()
		assert.Equal(t, c.UUID(), c.UUID())
	})

	t.Run("different seeds produce unique UUIDs", func(t *testing.T) {
		c := chaos.New(t.Name())
		results := make(map[uuid.UUID]bool)
		for i := 0; i < 1000; i++ {
			result := c.UUID()
			assert.False(t, results[result], "Expected unique result for each seed")
			results[result] = true
		}
		assert.Len(t, results, 1000, "Expected 1000 unique results")
	})

	t.Run("generates valid UUIDs", func(t *testing.T) {
		c := chaos.New(t.Name())
		for i := 0; i < 1000; i++ {
			id := c.UUID()
			assert.Equal(t, uuid.Version(4), id.Version())
			assert.Equal(t, uuid.RFC4122, id.Variant())
		}
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
