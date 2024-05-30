package helpers

import (
	"crypto/rand"
	"encoding/hex"
)

// GenerateRandomString generates a random string of the specified length.
func GenerateRandomString(length int) string {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err!= nil {
		panic(err)
	}
	return hex.EncodeToString(bytes)
}
