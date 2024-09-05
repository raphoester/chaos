package chaos

import (
	"testing"
	"time"

	"math"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInt(t *testing.T) {
	t.Run("deterministic output", func(t *testing.T) {
		assert.Equal(t, Int(100, "seed"), Int(100, "seed"))
	})

	t.Run("different seeds produce different results", func(t *testing.T) {
		assert.NotEqual(t, Int(100, "seed1"), Int(100, "seed2"))
	})

	t.Run("respects upper bound", func(t *testing.T) {
		for i := 0; i < 1000; i++ {
			result := Int(10, i)
			assert.GreaterOrEqual(t, result, 0)
			assert.LessOrEqual(t, result, 10)
		}
	})

	t.Run("edge case: zero", func(t *testing.T) {
		assert.Equal(t, 0, Int(0, "seed"))
	})

	t.Run("edge case: negative upper bound", func(t *testing.T) {
		result := Int(-10, "seed")
		assert.GreaterOrEqual(t, result, -10)
		assert.LessOrEqual(t, result, 0)
	})

	t.Run("edge case: MaxInt32", func(t *testing.T) {
		result := Int(math.MaxInt32, "seed")
		assert.GreaterOrEqual(t, result, 0)
		assert.LessOrEqual(t, result, math.MaxInt32)
	})
}

func TestBool(t *testing.T) {
	t.Run("deterministic output", func(t *testing.T) {
		assert.Equal(t, Bool("seed"), Bool("seed"))
	})

	t.Run("different seeds produce different results", func(t *testing.T) {
		found := false
		for i := 0; i < 100 && !found; i++ {
			if Bool(i) != Bool(i+1) {
				found = true
			}
		}
		assert.True(t, found, "Expected to find different boolean results")
	})
}

func TestDuration(t *testing.T) {
	t.Run("deterministic output", func(t *testing.T) {
		assert.Equal(t, Duration(time.Hour, "seed"), Duration(time.Hour, "seed"))
	})

	t.Run("respects upper bound", func(t *testing.T) {
		for i := 0; i < 1000; i++ {
			result := Duration(time.Hour, i)
			assert.GreaterOrEqual(t, result, time.Duration(0))
			assert.LessOrEqual(t, result, time.Hour)
		}
	})

	t.Run("edge case: zero duration", func(t *testing.T) {
		assert.Equal(t, time.Duration(0), Duration(0, "seed"))
	})

	t.Run("edge case: negative duration", func(t *testing.T) {
		result := Duration(-time.Hour, "seed")
		assert.GreaterOrEqual(t, result, -time.Hour)
		assert.LessOrEqual(t, result, time.Duration(0))
	})
}

func TestTime(t *testing.T) {
	t.Run("deterministic output", func(t *testing.T) {
		assert.Equal(t, Time("seed"), Time("seed"))
	})

	t.Run("time range", func(t *testing.T) {
		minTime := time.Unix(0, 0)
		maxTime := time.Unix(1<<32-1, 0)
		for i := 0; i < 1000; i++ {
			result := Time(i)
			assert.True(t, result.After(minTime) || result.Equal(minTime))
			assert.True(t, result.Before(maxTime) || result.Equal(maxTime))
		}
	})

	t.Run("edge case: far future time", func(t *testing.T) {
		farFuture := time.Now().AddDate(100, 0, 0)
		result := Time("farFuture")
		assert.True(t, result.Before(farFuture) || result.Equal(farFuture))
	})
}

func TestFloat(t *testing.T) {
	t.Run("deterministic output", func(t *testing.T) {
		assert.Equal(t, Float(1.0, "seed"), Float(1.0, "seed"))
	})

	t.Run("respects upper bound", func(t *testing.T) {
		for i := 0; i < 1000; i++ {
			result := Float(10.0, i)
			assert.GreaterOrEqual(t, result, 0.0)
			assert.Less(t, result, 10.0)
		}
	})

	t.Run("edge case: zero", func(t *testing.T) {
		assert.Equal(t, 0.0, Float(0.0, "seed"))
	})

	t.Run("edge case: negative upper bound", func(t *testing.T) {
		result := Float(-10.0, "seed")
		assert.GreaterOrEqual(t, result, -10.0)
		assert.Less(t, result, 0.0)
	})

	t.Run("edge case: very small positive number", func(t *testing.T) {
		result := Float(1e-10, "seed")
		assert.GreaterOrEqual(t, result, 0.0)
		assert.Less(t, result, 1e-10)
	})
}

func TestString(t *testing.T) {
	t.Run("deterministic output", func(t *testing.T) {
		assert.Equal(t, String(10, "seed"), String(10, "seed"))
	})

	t.Run("respects length", func(t *testing.T) {
		for i := 0; i < 100; i++ {
			result := String(i, "seed")
			assert.Len(t, result, i)
		}
	})

	t.Run("contains only alphanumeric characters", func(t *testing.T) {
		result := String(1000, "seed")
		for _, char := range result {
			assert.Contains(t, alphanumericalChars, string(char))
		}
	})

	t.Run("edge case: empty string", func(t *testing.T) {
		assert.Equal(t, "", String(0, "seed"))
	})
}

func TestSliceItem(t *testing.T) {
	t.Run("deterministic output", func(t *testing.T) {
		items := []int{1, 2, 3, 4, 5}
		assert.Equal(t, SliceItem(items, "seed"), SliceItem(items, "seed"))
	})

	t.Run("returns zero value for empty slice", func(t *testing.T) {
		var emptyIntSlice []int
		assert.Equal(t, 0, SliceItem(emptyIntSlice, "seed"))

		var emptyStringSlice []string
		assert.Equal(t, "", SliceItem(emptyStringSlice, "seed"))
	})

	t.Run("selects from all items", func(t *testing.T) {
		items := []int{1, 2, 3, 4, 5}
		selected := make(map[int]bool)
		for i := 0; i < 1000; i++ {
			selected[SliceItem(items, i)] = true
		}
		assert.Len(t, selected, len(items))
	})

	t.Run("edge case: single item slice", func(t *testing.T) {
		items := []int{42}
		assert.Equal(t, 42, SliceItem(items, "seed"))
	})
}

func TestUniqueSliceItems(t *testing.T) {
	t.Run("deterministic output", func(t *testing.T) {
		items := []int{1, 2, 3, 4, 5}
		result1, err1 := UniqueSliceItems(items, 3, "seed")
		result2, err2 := UniqueSliceItems(items, 3, "seed")
		require.NoError(t, err1)
		require.NoError(t, err2)
		assert.Equal(t, result1, result2)
	})

	t.Run("returns error when not enough items", func(t *testing.T) {
		items := []int{1, 2, 3}
		_, err := UniqueSliceItems(items, 4, "seed")
		assert.Error(t, err)
	})

	t.Run("returns unique items", func(t *testing.T) {
		items := []int{1, 2, 3, 4, 5}
		result, err := UniqueSliceItems(items, 3, "seed")
		require.NoError(t, err)
		assert.Len(t, result, 3)
		assert.Len(t, uniqueInts(result), 3)
	})

	t.Run("edge case: request all items", func(t *testing.T) {
		items := []int{1, 2, 3, 4, 5}
		result, err := UniqueSliceItems(items, 5, "seed")
		require.NoError(t, err)
		assert.ElementsMatch(t, items, result)
	})
}

func TestMustUniqueSliceItems(t *testing.T) {
	t.Run("returns unique items", func(t *testing.T) {
		items := []int{1, 2, 3, 4, 5}
		result := MustUniqueSliceItems(items, 3, "seed")
		assert.Len(t, result, 3)
		assert.Len(t, uniqueInts(result), 3)
	})

	t.Run("panics when not enough items", func(t *testing.T) {
		items := []int{1, 2, 3}
		assert.Panics(t, func() {
			MustUniqueSliceItems(items, 4, "seed")
		})
	})
}

func TestIntSlice(t *testing.T) {
	t.Run("deterministic output", func(t *testing.T) {
		assert.Equal(t, IntSlice(10, 5, "seed"), IntSlice(10, 5, "seed"))
	})

	t.Run("respects length", func(t *testing.T) {
		result := IntSlice(10, 5, "seed")
		assert.Len(t, result, 5)
	})

	t.Run("respects upper bound", func(t *testing.T) {
		result := IntSlice(10, 1000, "seed")
		for _, num := range result {
			assert.GreaterOrEqual(t, num, 0)
			assert.LessOrEqual(t, num, 10)
		}
	})

	t.Run("edge case: zero length", func(t *testing.T) {
		assert.Empty(t, IntSlice(10, 0, "seed"))
	})
}

func TestUUID(t *testing.T) {
	t.Run("deterministic output", func(t *testing.T) {
		assert.Equal(t, UUID("seed"), UUID("seed"))
	})

	t.Run("different seeds produce different UUIDs", func(t *testing.T) {
		assert.NotEqual(t, UUID("seed1"), UUID("seed2"))
	})

	t.Run("generates valid UUIDs", func(t *testing.T) {
		for i := 0; i < 1000; i++ {
			id := UUID(i)
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
