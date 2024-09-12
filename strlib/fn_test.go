package strlib_test

import (
	"github.com/tiyee/gokit/internal/assert"
	"github.com/tiyee/gokit/strlib"
	"testing"
)

func TestTrim(t *testing.T) {
	t.Parallel()

	as := assert.NewAssert(t, "TestTrim")

	str1 := "$ ab	cd $ "
	assert.Equal(as, "$ ab	cd $", strlib.Trim(str1))
	assert.Equal(as, "ab	cd", strlib.Trim(str1, "$"))
	assert.Equal(as, "abcd", strlib.Trim("\nabcd"))
}
func TestReverse(t *testing.T) {
	t.Parallel()

	as := assert.NewAssert(t, "TestReverse")

	assert.Equal(as, "cba", strlib.Reverse("abc"))
	assert.Equal(as, "54321", strlib.Reverse("12345"))
}
