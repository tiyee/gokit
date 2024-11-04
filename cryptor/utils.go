package cryptor

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"github.com/tiyee/gokit/slice"
)

func isOneOf[T comparable](want T, arr []T) bool {
	return slice.OneOf(want, arr)
}
func generateAesKey(key []byte, size int) []byte {
	genKey := make([]byte, size)
	copy(genKey, key)
	for i := size; i < len(key); {
		for j := 0; j < size && i < len(key); j, i = j+1, i+1 {
			genKey[j] ^= key[i]
		}
	}
	return genKey
}

func generateDesKey(key []byte) []byte {
	genKey := make([]byte, 8)
	copy(genKey, key)
	for i := 8; i < len(key); {
		for j := 0; j < 8 && i < len(key); j, i = j+1, i+1 {
			genKey[j] ^= key[i]
		}
	}
	return genKey
}

func pkcs7Padding(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padText...)
}

func pkcs7UnPadding(src []byte) []byte {
	length := len(src)
	unPadding := int(src[length-1])
	return src[:(length - unPadding)]
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
