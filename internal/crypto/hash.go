package crypto

import (
    "crypto/sha256"
    "encoding/hex"
)

// Hash computes SHA-256 of input and returns hex string
func Hash(data []byte) string {
    h := sha256.Sum256(data)
    return hex.EncodeToString(h[:])
}
