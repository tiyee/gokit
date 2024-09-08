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
