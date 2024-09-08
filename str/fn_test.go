package str_test

import (
	"github.com/tiyee/gokit/assert"
	"github.com/tiyee/gokit/str"
	"testing"
)

func TestTrim(t *testing.T) {
	t.Parallel()

	as := assert.NewAssert(t, "TestTrim")

	str1 := "$ ab	cd $ "

	assert.Equal(as, "$ ab	cd $", str.Trim(str1))
	assert.Equal(as, "ab	cd", str.Trim(str1, "$"))
	assert.Equal(as, "abcd", str.Trim("\nabcd"))
}
