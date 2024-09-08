package fn_test

import (
	"errors"
	"github.com/tiyee/gokit/assert"
	"github.com/tiyee/gokit/fn"
	"testing"
)

func TestMustOK(t *testing.T) {
	as := assert.NewAssert(t, "TestMustOK")
	assert.Panics(as, func() {
		fn.MustOK(errors.New("test"))
	})
}
