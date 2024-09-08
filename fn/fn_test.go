package fn_test

import (
	"errors"
	"github.com/tiyee/gokit/fn"
	"github.com/tiyee/gokit/internal/assert"
	"testing"
)

func TestMustOK(t *testing.T) {
	as := assert.NewAssert(t, "TestMustOK")
	assert.Panics(as, func() {
		fn.MustOK(errors.New("test"))
	})
}
