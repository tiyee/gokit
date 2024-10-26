package boollib

import "github.com/tiyee/gokit/constraints"

func ToInteger[T constraints.Integer](b bool) T {
	if b {
		return 1
	} else {
		return 0
	}
}
