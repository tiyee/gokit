package ptr

import "reflect"

// Deprecated:
func CopyPoint[T any](m T) T {
	vt := reflect.TypeOf(m).Elem()
	newV := reflect.New(vt)
	newV.Elem().Set(reflect.ValueOf(m).Elem())
	return newV.Interface().(T)
}
