package mathlib_test

import (
	"github.com/tiyee/gokit/internal/assert"
	"github.com/tiyee/gokit/mathlib"
	"testing"
)

func TestToString(t *testing.T) {
	as := assert.NewAssert(t, "TestToString")

	assert.Equal(as, "123", mathlib.ToString(123))
	assert.Equal(as, "-123", mathlib.ToString(-123))
}
