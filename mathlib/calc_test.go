package mathlib_test

import (
	"github.com/tiyee/gokit/internal/assert"
	"github.com/tiyee/gokit/mathlib"
	"testing"
)

func TestAbs(t *testing.T) {
	as := assert.NewAssert(t, "TestAbs")

	assert.Equal(as, int(1), mathlib.Abs(-1))
}
func TestCeil(t *testing.T) {
	as := assert.NewAssert(t, "TestCeil")
	total := 101
	pageSize := 20
	assert.LTE(as, 5, mathlib.Ceil[int](total, pageSize))
}
func TestFloor(t *testing.T) {
	as := assert.NewAssert(t, "TestFloor")
	total := 101
	pageSize := 20
	assert.GTE(as, 5, mathlib.Floor[int](total, pageSize))
}
func TestRound(t *testing.T) {
	as := assert.NewAssert(t, "TestRound")
	a := 4
	b := 5
	c := 3
	assert.Equal(as, 1, mathlib.Round(a, c))
	assert.LT(as, 1, mathlib.Round(b, c))

}
