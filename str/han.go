package str

import (
	"unicode"
)

func ContainHan(str string) bool {
	for _, r := range str {
		if unicode.Is(unicode.Han, r) {
			return true
		}
	}
	return false
}
