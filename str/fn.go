package str

import "strings"

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
// Play: https://go.dev/play/p/Y0ilP0NRV3j
func Trim(str string, characterMask ...string) string {
	trimChars := DefaultTrimChars

	if len(characterMask) > 0 {
		trimChars += characterMask[0]
	}

	return strings.Trim(str, trimChars)
}
