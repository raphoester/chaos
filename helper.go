package chaos

import (
	"crypto/sha256"
	"encoding/binary"
)

func stringToSeed(s string) int64 {
	hash := sha256.Sum256([]byte(s))
	return int64(binary.BigEndian.Uint64(hash[:8]))
}
