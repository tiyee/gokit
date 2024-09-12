// Package encodes provide some util for encode/decode data
package encodes

import (
	"encoding/base32"
	"encoding/base64"
)

// BaseEncoder interface
type BaseEncoder interface {
	Encode(dst []byte, src []byte)
	EncodeToString(src []byte) string
	Decode(dst []byte, src []byte) (n int, err error)
	DecodeString(s string) ([]byte, error)
	EncodedLen(n int) int
	DecodedLen(n int) int
}

//
// -------------------- base encode --------------------
//

// base32 encoding with no padding
type Encoding string

const (
	B32Std Encoding = "B32Std"

	B32Hex Encoding = "B32Hex"
	B64Std Encoding = "B64Std"
	B64URL Encoding = "B64URL"
)

var (
	b32Std = base32.StdEncoding.WithPadding(base32.NoPadding)
	b32Hex = base32.HexEncoding.WithPadding(base32.NoPadding)
)

// base64 encoding with no padding
var (
	b64Std = base64.StdEncoding.WithPadding(base64.NoPadding)
	b64URL = base64.URLEncoding.WithPadding(base64.NoPadding)
)

func getEncoder(coding Encoding) BaseEncoder {
	switch coding {
	case B32Std:
		return b32Std
	case B32Hex:
		return b32Hex
	case B64Std:
		return b64Std
	case B64URL:
		return b64URL
	default:
		panic("unknown encoding")
	}
}
func EncodeString(coding Encoding, str string) string {
	return getEncoder(coding).EncodeToString([]byte(str))
}
func Encode(coding Encoding, src []byte) []byte {
	encoder := getEncoder(coding)
	buf := make([]byte, encoder.EncodedLen(len(src)))
	encoder.Encode(buf, src)
	return buf
}
func Decode(coding Encoding, str []byte) []byte {
	encoder := getEncoder(coding)
	dbuf := make([]byte, encoder.DecodedLen(len(str)))
	n, _ := encoder.Decode(dbuf, str)
	return dbuf[:n]
}
func DecodeString(coding Encoding, str string) string {
	encoder := getEncoder(coding)
	dec, _ := encoder.DecodeString(str)
	return string(dec)
}
func Base64URLEncoding(bs []byte) string {
	return getEncoder(B64URL).EncodeToString(bs)
}
func Base64StdEncoding(bs []byte) string {
	return getEncoder(B64Std).EncodeToString(bs)
}
func Base64URLDecoding(s string) ([]byte, error) {
	return getEncoder(B64URL).DecodeString(s)
}
func Base64StdDecoding(s string) ([]byte, error) {
	return getEncoder(B64Std).DecodeString(s)
}
