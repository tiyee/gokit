package slice

import (
	"cmp"
	"sort"
)

func Last[E any](s []E) (E, bool) {
	if len(s) == 0 {
		var zero E
		return zero, false
	}
	return s[len(s)-1], true
}
func Keys[K comparable, V any](m map[K]V) []K {
	if m == nil {
		return []K{}
	}
	arr := make([]K, 0, len(m))
	for k, _ := range m {
		arr = append(arr, k)
	}
	return arr
}

// Equal compare two slice
// @notice: don't check sequence
func Equal[T comparable](a, b []T) bool {
	if a == nil || b == nil {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	return equal(a, b)
}
func equal[T comparable](a, b []T) bool {
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
func StrictEqual[T cmp.Ordered](a, b []T) bool {
	if a == nil || b == nil {
		return false
	}
	if len(a) != len(b) {
		return false
	}

	sort.Sort(Slice[T](a))
	sort.Sort(Slice[T](b))
	return equal[T](a, b)
}
func Merge[T any](a []T, b []T, cc ...[]T) []T {
	if a == nil || b == nil {
		return []T{}
	}
	n := len(a) + len(b)
	for _, c := range cc {
		n += len(c)
	}
	result := make([]T, 0, n)
	result = append(result, a...)
	result = append(result, b...)
	for _, c := range cc {
		result = append(result, c...)
	}
	return result

}
