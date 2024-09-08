package hashlib

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"github.com/tiyee/gokit/encodes"
	"hash"
	"hash/crc32"
	"hash/crc64"
	"io"
)

// HashType hash algorithm names
type HashType string

const (
	AlgoCRC32      HashType = "crc32"
	AlgoCRC64      HashType = "crc64"
	AlgoMD5        HashType = "md5"
	AlgoSHA1       HashType = "sha1"
	AlgoSHA224     HashType = "sha224"
	AlgoSHA256     HashType = "sha256"
	AlgoSHA384     HashType = "sha384"
	AlgoSHA512     HashType = "sha512"
	AlgoSHA512_224 HashType = "sha512_224"
	AlgoSHA512_256 HashType = "sha512_256"
)

// MD5 generate md5 string by given src
func MD5(src []byte) string {
	return string(HexBytes(AlgoMD5, src))
}

// ShortMD5 Generate a 16-bit md5 bytes.
// remove first 8 and last 8 bytes from 32-bit md5.
func ShortMD5(src []byte) string {
	return string(HexBytes(AlgoMD5, src)[8:24])
}

// FileHash Generate a 64-bit md5 bytes from a file or file-like io
func FileHash(algo HashType, f io.Reader) (string, error) {
	if bs, err := HashSumReader(algo, f); err == nil {
		return string(bs), nil
	} else {
		return "", err
	}
}

// Hash generate hex hash string by given algorithm
func Hash(algo HashType, src []byte) string {
	return string(HexBytes(algo, src))
}

// HexBytes generate hex hash bytes by given algorithm
func HexBytes(algo HashType, src []byte) []byte {
	bs := HashSum(algo, src)
	dst := make([]byte, hex.EncodedLen(len(bs)))
	hex.Encode(dst, bs)
	return dst
}

// Hash32 generate hash by given algorithm, then use base32 encode.
func Hash32(algo HashType, src []byte) string {
	return string(Base32Bytes(algo, src))
}

// Base32Bytes generate base32 hash bytes by given algorithm
func Base32Bytes(algo HashType, src []byte) []byte {
	bs := HashSum(algo, src)
	return encodes.Encode("B32Hex", bs)

}

// Hash64 generate hash by given algorithm, then use base64 encode.
func Hash64(algo HashType, src []byte) string {
	return string(Base64Bytes(algo, src))
}

// Base64Bytes generate base64 hash bytes by given algorithm
func Base64Bytes(algo HashType, src []byte) []byte {
	bs := HashSum(algo, src)
	return encodes.Encode("B64Std", bs)

}

// HashSum generate hash sum bytes by given algorithm
func HashSum(algo HashType, src []byte) []byte {
	hh := NewHash(algo)

	hh.Write(src)
	return hh.Sum(nil)
}

// HashSumReader generate file or file-like hash sum bytes by given algorithm
func HashSumReader(algo HashType, src io.Reader) ([]byte, error) {
	hh := NewHash(algo)
	if _, err := io.Copy(hh, src); err != nil {
		return nil, err
	}
	return hh.Sum(nil), nil
}

// NewHash create hash.Hash instance
//
// algo: crc32, crc64, md5, sha1, sha224, sha256, sha384, sha512, sha512_224, sha512_256
func NewHash(algo HashType) hash.Hash {
	switch algo {
	case AlgoCRC32:
		return crc32.NewIEEE()
	case AlgoCRC64:
		return crc64.New(crc64.MakeTable(crc64.ISO))
	case AlgoMD5:
		return md5.New()
	case AlgoSHA1:
		return sha1.New()
	case AlgoSHA224:
		return sha256.New224()
	case AlgoSHA256:
		return sha256.New()
	case AlgoSHA384:
		return sha512.New384()
	case AlgoSHA512:
		return sha512.New()
	case AlgoSHA512_224:
		return sha512.New512_224()
	case AlgoSHA512_256:
		return sha512.New512_256()
	default:
		panic("invalid hash algorithm:" + algo)
	}
}
