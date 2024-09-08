package mathlib

import (
	"github.com/tiyee/gokit/constraints"
	"strconv"
)

// ToString convert a integer to string with base 10
func ToString[T constraints.Integer](input T) string {
	return strconv.FormatInt(int64(input), 10)
}
