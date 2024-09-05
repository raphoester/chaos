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
	return singleton.Int32(n + 1)
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

// Time returns a random time between Unix epoch and 2106-02-07 08:28:16.
func Time() time.Time {
	return singleton.Time()
}

// Time returns a deterministic time between Unix epoch and 2106-02-07 08:28:16.
func (c *Chaos) Time() time.Time {
	return time.Unix(c.Int64(1<<32), 0)
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

// Float64 returns a random float64 between 0 and n.
func Float64(n float64) float64 {
	return singleton.Float64(n)
}

// Float64 returns a deterministic float64 between 0 and n.
func (c *Chaos) Float64(n float64) float64 {
	r := c.rand()
	return r.Float64() * n
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
