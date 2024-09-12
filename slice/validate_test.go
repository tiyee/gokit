package slice_test

import (
	"github.com/tiyee/gokit/internal/assert"
	"github.com/tiyee/gokit/slice"
	"testing"
)

func TestOneOf(t *testing.T) {
	as := assert.NewAssert(t, "TestOf")
	assert.Equal(as, true, slice.OneOf(1, []int{1, 2, 3, 4}))
}
func TestNotOneOf(t *testing.T) {
	as := assert.NewAssert(t, "TestNotOneOf")
	assert.Equal(as, true, slice.NotIn(0, []int{1, 2, 3}))
}
