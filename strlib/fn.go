package strlib

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// DefaultTrimChars are the characters which are stripped by Trim* functions in default.
var DefaultTrimChars = string([]byte{
	'\t', // Tab.
	'\v', // Vertical tab.
	'\n', // New line (line feed).
	'\r', // Carriage return.
	'\f', // New page.
	' ',  // Ordinary space.
	0x00, // NUL-byte.
	0x85, // Delete.
	0xA0, // Non-breaking space.
})

// Trim strips whitespace (or other characters) from the beginning and end of a string.
// The optional parameter `characterMask` specifies the additional stripped characters.
func Trim(str string, characterMask ...string) string {
	trimChars := DefaultTrimChars

	if len(characterMask) > 0 {
		trimChars += characterMask[0]
	}

	return strings.Trim(str, trimChars)
}

// Reverse returns string whose char order is reversed to the given string.
func Reverse(s string) string {
	l := len(s)
	if l < 2 {
		return s
	}
	r := []rune(s)
	for i, j := 0, l-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
func NonceStr() string {
	ns := time.Now().UnixNano()
	bs := md5.Sum([]byte(strconv.FormatInt(ns, 10)))
	return fmt.Sprintf("%x", bs)
}
