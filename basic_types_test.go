package chaos_test

import (
	"fmt"
	"testing"
	"time"

	"math"

	"github.com/google/uuid"
	"github.com/raphoester/chaos"
	"github.com/stretchr/testify/assert"
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
