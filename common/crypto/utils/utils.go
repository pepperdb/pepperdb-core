package utils

import (
	"crypto/rand"
	"io"
)

// RandomCSPRNG a cryptographically secure pseudo-random number generator
func RandomCSPRNG(n int) []byte {
	buff := make([]byte, n)
	_, err := io.ReadFull(rand.Reader, buff)
	if err != nil {
		panic("reading from crypto/rand failed: " + err.Error())
	}
	return buff
}

// ZeroBytes clears byte slice.
func ZeroBytes(bytes []byte) {
	for i := range bytes {
		bytes[i] = 0
	}
}
