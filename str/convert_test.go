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
