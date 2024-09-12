package ptrlib

import "github.com/tiyee/gokit/constraints"

func ToValue[T constraints.Ordered](v *T, missing T) T {
	if v == nil {
		return missing
	}
	return *v
}
