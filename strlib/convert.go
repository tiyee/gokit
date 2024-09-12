package strlib

import (
	"github.com/tiyee/gokit/constraints"
	"strconv"
	"strings"
)

func StringOrEmpty(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}
func ToInteger[T constraints.Integer](s string, missing T) T {
	if s == "" {
		return missing
	}
	if n, err := strconv.ParseInt(s, 10, 64); err == nil {
		return T(n)
	} else {
		return missing
	}
}
func ToFloat[T constraints.Float](s string, missing T) T {
	if s == "" {
		return missing
	}
	if n, err := strconv.ParseFloat(s, 64); err == nil {
		return T(n)
	} else {
		return missing
	}
}
func ToBool(s string, fn ...func(s string) bool) bool {
	if s == "" {
		return false
	}
	switch strings.ToUpper(s) {
	case "TRUE":
		return true
	case "FALSE":
		return false
	case "1":
		return true
	case "0":
		return false
	case "1.0":
		return true
	case "0.0":
		return false

	default:
		for _, f := range fn {
			if f(s) {
				return true
			}
		}
		return false

	}
}
