package str_test

import (
	"github.com/tiyee/gokit/internal/assert"
	"github.com/tiyee/gokit/str"
	"testing"
)

func TestContainHan(t *testing.T) {
	as := assert.NewAssert(t, "TestOf")
	assert.Equal(as, true, str.ContainHan("121我爱"))
	assert.Equal(as, false, str.ContainHan("121ssssssssssss"))
}
