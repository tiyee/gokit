package helps

import (
	"reflect"
)

func CopyPoint[T any](m T) T {
	vt := reflect.TypeOf(m).Elem()
	newoby := reflect.New(vt)
	newoby.Elem().Set(reflect.ValueOf(m).Elem())
	return newoby.Interface().(T)
}
