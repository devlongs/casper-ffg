package utils

import (
	"crypto/sha256"
)

func CalculateHash(data []byte) []byte {
	hash := sha256.Sum256(data)
	return hash[:]
}
