package str_test

import (
	"github.com/tiyee/gokit/internal/assert"
	"github.com/tiyee/gokit/str"
	"testing"
)

func TestStringOrEmpty(t *testing.T) {
	as := assert.NewAssert(t, "TestStringOrEmpty")
	s := "12a"
	assert.Equal(as, s, str.StringOrEmpty(&s))
	assert.Equal(as, "", str.StringOrEmpty(nil))
}
func TestToInteger(t *testing.T) {
	as := assert.NewAssert(t, "TestToInteger")
	s := "1"
	assert.Equal(as, int8(1), str.ToInteger[int8](s, 0))
}
func TestToFloat(t *testing.T) {
	as := assert.NewAssert(t, "TestToFloat")
	s := "1.23"
	assert.Equal(as, float32(1.23), str.ToFloat[float32](s, 0))
}
func TestToBool(t *testing.T) {
	as := assert.NewAssert(t, "TestToBool")
	s := "true"
	assert.Equal(as, true, str.ToBool(s))
}
