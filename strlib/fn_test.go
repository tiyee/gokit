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
func TestStrip(t *testing.T) {
	t.Parallel()

	as := assert.NewAssert(t, "TestStrip")

	str1 := "$ ab	cd $ "
	assert.Equal(as, "$ ab	cd $", strlib.Strip(str1))
	assert.Equal(as, "ab	cd", strlib.Strip(str1, "$"))
	assert.Equal(as, "abcd", strlib.Strip("\nabcd"))
}
func TestLStrip(t *testing.T) {
	t.Parallel()

	as := assert.NewAssert(t, "TestLStrip")

	str1 := "$ ab	cd $ "
	assert.Equal(as, "$ ab\tcd $ ", strlib.LStrip(str1))
	assert.Equal(as, "ab	cd $ ", strlib.LStrip(str1, "$"))
	assert.Equal(as, "abcd", strlib.LStrip("\nabcd"))
}
func TestRStrip(t *testing.T) {
	t.Parallel()

	as := assert.NewAssert(t, "TestRStrip")

	str1 := "$ ab	cd $ "
	assert.Equal(as, "$ ab\tcd $", strlib.RStrip(str1))
	assert.Equal(as, "$ ab\tcd", strlib.RStrip(str1, "$"))
	assert.Equal(as, "\nabcd", strlib.RStrip("\nabcd"))
}
func TestReverse(t *testing.T) {
	t.Parallel()

	as := assert.NewAssert(t, "TestReverse")

	assert.Equal(as, "cba", strlib.Reverse("abc"))
	assert.Equal(as, "54321", strlib.Reverse("12345"))
}
