package chaos

import (
	"time"

	"github.com/google/uuid"
)

// Int32 generates a deterministic integer between 0 and n (inclusive) based on the provided seed.
// Behavior:
//   - If n <= 0, the function returns 0.
//   - The generated integer is guaranteed to be within the range [0, n],
//     including both 0 and n as possible values.
func Int32(n int32) int32 {
	return singleton.Int32(n)
}

// Int32 generates a deterministic integer between 0 and n (inclusive) based on the provided seed.
// Behavior:
//   - If n <= 0, the function returns 0.
//   - The generated integer is guaranteed to be within the range [0, n],
//     including both 0 and n as possible values.
func (c *Chaos) Int32(n int32) int32 {
	if n <= 0 {
		return 0
	}
	r := c.rand()
	return r.Int31n(n + 1)
}

// Int32Between generates a deterministic integer between min and max (inclusive) based on the provided seed.
// Behavior:
//   - If min > max, the values are swapped.
//   - The generated integer is guaranteed to be within the range [min, max],
//     including both min and max as possible values.
func Int32Between(min, max int32) int32 {
	return singleton.Int32Between(min, max)
}

// Int32Between generates a deterministic integer between min and max (inclusive) based on the provided seed.
// Behavior:
//   - If min > max, the values are swapped.
//   - The generated integer is guaranteed to be within the range [min, max],
//     including both min and max as possible values.
func (c *Chaos) Int32Between(min, max int32) int32 {
	if min > max {
		min, max = max, min
	}
	return c.Int32(max-min) + min
}

// Int64 generates a deterministic integer between 0 and n (inclusive) based on the provided seed.
// Behavior:
//   - If n <= 0, the function returns 0.
//   - The generated integer is guaranteed to be within the range [0, n],
//     including both 0 and n as possible values.
func Int64(n int64) int64 {
	return singleton.Int64(n)
}

// Int64 generates a deterministic integer between 0 and n (inclusive) based on the provided seed.
// Behavior:
//   - If n <= 0, the function returns 0.
//   - The generated integer is guaranteed to be within the range [0, n],
//     including both 0 and n as possible values.
func (c *Chaos) Int64(n int64) int64 {
	if n <= 0 {
		return 0
	}
	r := c.rand()
	return r.Int63n(n + 1)
}

// Int64Between generates a deterministic integer between min and max (inclusive) based on the provided seed.
// Behavior:
//   - If min > max, the values are swapped.
//   - The generated integer is guaranteed to be within the range [min, max],
//     including both min and max as possible values.
func Int64Between(min, max int64) int64 {
	return singleton.Int64Between(min, max)
}

// Int64Between generates a deterministic integer between min and max (inclusive) based on the provided seed.
// Behavior:
//   - If min > max, the values are swapped.
//   - The generated integer is guaranteed to be within the range [min, max],
//     including both min and max as possible values.
func (c *Chaos) Int64Between(min, max int64) int64 {
	if min > max {
		min, max = max, min
	}
	return c.Int64(max-min) + min
}

// Int generates a deterministic integer between 0 and n (inclusive) based on the provided seed.
// Behavior:
//   - If n <= 0, the function returns 0.
//   - The generated integer is guaranteed to be within the range [0, n],
//     including both 0 and n as possible values.
func Int(n int) int {
	return singleton.Int(n)
}

// Int generates a deterministic integer between 0 and n (inclusive) based on the provided seed.
// Behavior:
//   - If n <= 0, the function returns 0.
//   - The generated integer is guaranteed to be within the range [0, n],
//     including both 0 and n as possible values.
func (c *Chaos) Int(n int) int {
	if n <= 0 {
		return 0
	}
	r := c.rand()
	return r.Intn(n + 1)
}

// IntBetween generates a deterministic integer between min and max (inclusive) based on the provided seed.
// Behavior:
//   - If min > max, the values are swapped.
//   - The generated integer is guaranteed to be within the range [min, max],
//     including both min and max as possible values.
func IntBetween(min, max int) int {
	return singleton.IntBetween(min, max)
}

// IntBetween generates a deterministic integer between min and max (inclusive) based on the provided seed.
// Behavior:
//   - If min > max, the values are swapped.
//   - The generated integer is guaranteed to be within the range [min, max],
//     including both min and max as possible values.
func (c *Chaos) IntBetween(min, max int) int {
	if min > max {
		min, max = max, min
	}
	return c.Int(max-min) + min
}

// Bool generates a deterministic boolean.
func Bool() bool {
	return singleton.Bool()
}

// Bool generates a deterministic boolean.
func (c *Chaos) Bool() bool {
	return c.Int64(2)%2 == 0
}

// Duration returns a random duration between 0 and n.
func Duration(n time.Duration) time.Duration {
	return singleton.Duration(n)
}

// Duration returns a deterministic duration between 0 and n.
func (c *Chaos) Duration(n time.Duration) time.Duration {
	return time.Duration(c.Int64(n.Nanoseconds()))
}

// DurationBetween returns a random duration between min and max.
func DurationBetween(min, max time.Duration) time.Duration {
	return singleton.DurationBetween(min, max)
}

func (c *Chaos) DurationBetween(min, max time.Duration) time.Duration {
	return time.Duration(c.Int64Between(min.Nanoseconds(), max.Nanoseconds()))
}

// Time returns a random time between Unix epoch and 2106-02-07 08:28:16.
func Time() time.Time {
	return singleton.Time()
}

// Time returns a deterministic time between Unix epoch and 2106-02-07 08:28:16.
func (c *Chaos) Time() time.Time {
	return time.Unix(c.Int64(1<<32), 0)
}

// TimeBetween returns a random time between min and max.
func TimeBetween(min, max time.Time) time.Time {
	return singleton.TimeBetween(min, max)
}

// TimeBetween returns a deterministic time between min and max.
func (c *Chaos) TimeBetween(min, max time.Time) time.Time {
	minUnix := min.Unix()
	maxUnix := max.Unix()
	return time.Unix(c.Int64Between(minUnix, maxUnix), 0)
}

// Float32 returns a random float32 between 0 and n.
func Float32(n float32) float32 {
	return singleton.Float32(n)
}

// Float32 returns a deterministic float32 between 0 and n.
func (c *Chaos) Float32(n float32) float32 {
	r := c.rand()
	return r.Float32() * n
}

// Float32Between returns a random float32 between min and max.
func Float32Between(min, max float32) float32 {
	return singleton.Float32Between(min, max)
}

// Float32Between returns a deterministic float32 between min and max.
func (c *Chaos) Float32Between(min, max float32) float32 {
	return c.Float32(max-min) + min
}

// Float64 returns a random float64 between 0 and n.
func Float64(n float64) float64 {
	return singleton.Float64(n)
}

// Float64 returns a deterministic float64 between 0 and n.
func (c *Chaos) Float64(n float64) float64 {
	r := c.rand()
	return r.Float64() * n
}

// Float64Between returns a random float64 between min and max.
func Float64Between(min, max float64) float64 {
	return singleton.Float64Between(min, max)
}

// Float64Between returns a deterministic float64 between min and max.
func (c *Chaos) Float64Between(min, max float64) float64 {
	return c.Float64(max-min) + min
}

const alphanumericalChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// String returns a random string of <length> alphanumerical characters.
func String(length int) string {
	return singleton.String(length)
}

// String returns a random string of <length> alphanumerical characters.
func (c *Chaos) String(length int) string {
	ret := ""
	for i := 0; i < length; i++ {
		index := c.Int(len(alphanumericalChars) - 1)
		ret += string(alphanumericalChars[index])
	}
	return ret
}

// IntSlice returns a slice of random numbers between 0 and high included.
// The length of the slice is length.
func IntSlice(high int, length int) []int {
	return singleton.IntSlice(high, length)
}

// IntSlice returns a slice of deterministic numbers between 0 and high included.
// The length of the slice is length.
func (c *Chaos) IntSlice(high int, length int) []int {
	result := make([]int, length)
	for i := 0; i < length; i++ {
		result[i] = c.Int(high)
	}

	return result
}

func UUID() uuid.UUID {
	return singleton.UUID()
}

// UUID returns a random UUID.
func (c *Chaos) UUID() uuid.UUID {
	rnd := c.rand()
	id, err := uuid.NewRandomFromReader(rnd)
	if err != nil {
		panic(err)
	}
	return id
}
