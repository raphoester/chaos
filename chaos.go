package chaos

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"math/rand"
)

type Chaos struct {
	count int
	fixed bool
	seed  string
}

func New(seed string) *Chaos {
	return &Chaos{
		count: 0,
		fixed: false,
		seed:  seed,
	}
}

// Fix freezes the chaos.
// When chaos is fixed, the values generated are always the same.
// That means the same method will always return the same value.
func (c *Chaos) Fix() {
	c.fixed = true
}

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
