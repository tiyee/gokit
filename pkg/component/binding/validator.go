package binding

import (
	"errors"
	"github.com/tiyee/gokit/pkg/constraints"
)

type IField[T constraints.FieldType] interface {
	Value() T
	Valid() error
	Default(n T)
}

type IFn[T constraints.FieldType] interface {
}

type IntegerField[T constraints.Integer] struct {
	key        string
	value      T
	required   bool
	missing    T
	validators []func(s T) error
}

func Min[T constraints.Integer](min T) func(src T) error {
	return func(src T) error {
		if src < min {
			return errors.New("min error ")
		}
		return nil
	}

}
func Max[T constraints.Integer](max T) func(src T) error {
	return func(src T) error {
		if src > max {
			return errors.New("max error ")
		}
		return nil
	}

}
func OneOf[T constraints.Ordered](des []T) func(src T) error {
	return func(src T) error {
		for _, item := range des {
			if item == src {
				return nil
			}
		}
		return errors.New("oneOf error ")
	}

}
func Equal[T constraints.Ordered](src, des T) func(src T) error {
	return func(src T) error {
		if src != des {
			return errors.New("equal error ")
		}
		return nil
	}
}

func Integer[T constraints.Integer](key string, required bool, validators ...func(s T) error) *IntegerField[T] {
	return &IntegerField[T]{key: key, required: required, validators: validators}
}
func (it *IntegerField[T]) Valid() error {
	for _, fn := range it.validators {
		if err := fn(it.value); err != nil {
			return err
		}
	}
	return nil
}
func (it *IntegerField[T]) Value() T {
	return it.value
}
func (it *IntegerField[T]) Default(n T) {
	it.missing = n
}
func test() {

}
