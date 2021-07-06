package hasher

import (
	"crypto/sha256"
	"fmt"
	"strconv"
)

// Convert a string to a SHA-256 hash
func StringToSha256(input string) string {
	bytesSlice := []byte(input)
	return BytesSliceToSha256(bytesSlice)
}

// Convert an int to a SHA-256 hash
func IntToSha256(input int) string {
	bytesSlice := []byte(strconv.Itoa(input))
	return BytesSliceToSha256(bytesSlice)
}

// Convert a slice of bytes to a SHA-256 hash
func BytesSliceToSha256(input []byte) string {
	h := sha256.New()
	h.Write(input)

	return fmt.Sprintf("%x", h.Sum(nil))
}
