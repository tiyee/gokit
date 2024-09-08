package slice

import "cmp"

type Slice[T cmp.Ordered] []T

func (s Slice[T]) Len() int {
	return len(s)
}

func (s Slice[T]) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s Slice[T]) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
