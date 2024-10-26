package ptrlib

func ToValue[T any](v *T, missing T) T {
	if v == nil {
		return missing
	}
	return *v
}
