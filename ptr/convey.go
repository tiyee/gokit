package ptr

import "github.com/tiyee/gokit/internal/constraints"

func ToValue[T constraints.Ordered](v *T, missing T) T {
	if v == nil {
		return missing
	}
	return *v
}
