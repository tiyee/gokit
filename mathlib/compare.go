package mathlib

import "github.com/tiyee/gokit/internal/constraints"

func Max[T constraints.Ordered](i T, n ...T) T {
	for _, nn := range n {
		if nn > i {
			i = nn
		}
	}
	return i
}
func Min[T constraints.Ordered](i T, n ...T) T {
	for _, nn := range n {
		if nn < i {
			i = nn
		}
	}
	return i
}

// InRange check if val in int/float range [min, max]
func InRange[T constraints.Ordered](val, min, max T) bool {
	return val >= min && val <= max
}

// OutRange check if val not in int/float range [min, max]
func OutRange[T constraints.Ordered](val, min, max T) bool {
	return val < min || val > max
}
