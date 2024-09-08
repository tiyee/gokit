package ptr_test

import (
	"github.com/tiyee/gokit/internal/assert"
	"github.com/tiyee/gokit/ptr"
	"testing"
)

func TestToValue(t *testing.T) {
	as := assert.NewAssert(t, "TestToValue")
	a := 1
	b := &a
	assert.Equal(as, a, ptr.ToValue(b, 1))
	assert.Equal(as, 1, ptr.ToValue(nil, 1))
}
