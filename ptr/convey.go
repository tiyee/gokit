package ptr

import "golang.org/x/exp/constraints"

func ToValue[T constraints.Ordered](v *T, missing T) T {
	if v == nil {
		return missing
	}
	return *v
}
