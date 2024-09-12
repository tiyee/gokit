package strlib_test

import (
	"github.com/tiyee/gokit/internal/assert"
	"github.com/tiyee/gokit/strlib"
	"testing"
)

func TestStringOrEmpty(t *testing.T) {
	as := assert.NewAssert(t, "TestStringOrEmpty")
	s := "12a"
	assert.Equal(as, s, strlib.StringOrEmpty(&s))
	assert.Equal(as, "", strlib.StringOrEmpty(nil))
}
func TestToInteger(t *testing.T) {
	as := assert.NewAssert(t, "TestToInteger")
	s := "1"
	assert.Equal(as, int8(1), strlib.ToInteger[int8](s, 0))
}
func TestToFloat(t *testing.T) {
	as := assert.NewAssert(t, "TestToFloat")
	s := "1.23"
	assert.Equal(as, float32(1.23), strlib.ToFloat[float32](s, 0))
}
func TestToBool(t *testing.T) {
	as := assert.NewAssert(t, "TestToBool")
	s := "true"
	assert.Equal(as, true, strlib.ToBool(s))
}
