package boollib_test

import (
	"github.com/tiyee/gokit/boollib"
	"github.com/tiyee/gokit/internal/assert"
	"testing"
)

func TestToInteger(t *testing.T) {
	as := assert.NewAssert(t, "TestToInteger")
	assert.Equal(as, 1, boollib.ToInteger[int](true))
	assert.Equal(as, 0, boollib.ToInteger[int](false))
	assert.Equal(as, int8(1), boollib.ToInteger[int8](true))
}
