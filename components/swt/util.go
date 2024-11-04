package swt

import (
	"crypto/rand"
	"encoding/base64"
	"strings"
)

// EncodeSegment Encode JWT specific base64url encoding with padding stripped
func EncodeSegment(seg []byte) string {
	return strings.TrimRight(base64.URLEncoding.EncodeToString(seg), "=")
}

// DecodeSegment Decode JWT specific base64url encoding with padding stripped
func DecodeSegment(seg string) ([]byte, error) {
	if l := len(seg) % 4; l > 0 {
		seg += strings.Repeat("=", 4-l)
	}

	return base64.URLEncoding.DecodeString(seg)
}

// GenerateRandomString generates a random string of the specified length
func GenerateRandomString(length int) string {
	return base64.URLEncoding.EncodeToString(GenerateRandomKey(length))
}

// GenerateRandomKey generates a random bytes of the specified length
func GenerateRandomKey(length int) []byte {
	b := make([]byte, length)
	if _, err := rand.Read(b); err == nil {
		return b
	} else {
		return b
	}
}
