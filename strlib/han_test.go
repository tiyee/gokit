package strlib_test

import (
	"github.com/tiyee/gokit/internal/assert"
	"github.com/tiyee/gokit/strlib"
	"testing"
)

func TestContainHan(t *testing.T) {
	as := assert.NewAssert(t, "TestOf")
	assert.Equal(as, true, strlib.ContainHan("121我爱"))
	assert.Equal(as, false, strlib.ContainHan("121ssssssssssss"))
}
