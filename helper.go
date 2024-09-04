package chaos

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
)

func stringToSeed(s []any) int64 {
	hash := sha256.Sum256([]byte(fmt.Sprint(s...)))
	return int64(binary.BigEndian.Uint64(hash[:8]))
}
