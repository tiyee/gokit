package slice

func OneOf[T comparable](input T, candidates []T) bool {
	for _, item := range candidates {
		if item == input {
			return true
		}
	}
	return false
}
func NotIn[T comparable](input T, candidates []T) bool {
	return !OneOf[T](input, candidates)
}
