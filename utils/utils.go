package utils

import (
	"crypto/sha256"
	"encoding/binary"
)

func CalculateHash(data []byte) []byte {
	hash := sha256.Sum256(data)
	return hash[:]
}

func IntToBytes(num int) []byte {
	bytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(bytes, uint32(num))
	return bytes
}
