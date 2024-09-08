package mathlib

import (
	"golang.org/x/exp/constraints"
	"math"
)

type Numeric interface {
	constraints.Integer | constraints.Float
}

func Abs[T Numeric](n T) T {
	if n < 0 {
		return -n
	}
	return n
}

func Round[T Numeric](n T) T {
	return T(math.Round(float64(n)))
}
func Ceil[T Numeric](a, b T) T {
	if b == 0 {
		return 0
	}
	return T(math.Ceil(float64(a) / float64(b)))
}
func Floor[T Numeric](a, b T) T {
	if b == 0 {
		return 0
	}
	return T(math.Floor(float64(a) / float64(b)))
}
func Sum[T Numeric](n []T, missing T) T {

	for _, v := range n {
		missing += v
	}
	return missing
}

// Average return average value of numbers.
func Average[T Numeric](numbers ...T) T {

	n := len(numbers)
	if n == 0 {
		return 0
	}
	sum := Sum(numbers, 0)
	return sum / T(n)
}
func Percent[T Numeric](a, b T) float64 {
	if b == 0 {
		return 0
	}
	return float64(a) / float64(b)
}
