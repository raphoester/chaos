package chaos

import (
	"crypto/sha256"
	"encoding/binary"
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

// MustIntInRange returns a deterministic integer between min and max (inclusive).
// It panics if max is less than min.
func MustIntInRange(min, max int, seed ...any) int {
	if max < min {
		panic("max must be greater than or equal to min")
	}
	return min + Int(max-min, seed...)
}

// Bool returns a deterministic boolean.
func Bool(seed ...any) bool {
	return stringToSeed(seed)%2 == 0
}

// Duration returns a deterministic duration between 0 and n.
func Duration(n time.Duration, seed ...any) time.Duration {
	return time.Duration(Int(n.Nanoseconds(), seed))
}

// MustDurationInRange returns a deterministic duration between min and max.
// It panics if max is less than min.
func MustDurationInRange(min, max time.Duration, seed ...any) time.Duration {
	if max < min {
		panic("max must be greater than or equal to min")
	}
	return min + Duration(max-min, seed...)
}

// Fix freezes the chaos.
// When chaos is fixed, the values generated are always the same.
// That means the same method will always return the same value.
func (c *Chaos) Fix() {
	c.fixed = true
}

// Unfix un-freezes the chaos.
// This results in new values being generated for each method call.
func (c *Chaos) Unfix() {
	c.fixed = false
}

func (c *Chaos) rand() *rand.Rand {
	if !c.fixed {
		c.count++
	}
	seed := fmt.Sprintf("%s-%d", c.seed, c.count)
	hash := sha256.Sum256([]byte(seed))
	intSeed := int64(binary.BigEndian.Uint64(hash[:8]))
	return rand.New(rand.NewSource(intSeed))
}
