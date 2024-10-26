package boollib_test

import (
	"errors"
	"fmt"
	"github.com/tiyee/gokit/boollib"
	"github.com/tiyee/gokit/internal/assert"
	"testing"
)

func TestCallOn(t *testing.T) {
	okFunc := func() error {
		fmt.Println("ok")
		return nil
	}
	errFunc := func() error {
		fmt.Println("err")
		return errors.New("err")
	}
	as := assert.NewAssert(t, "TestCallOn")
	assert.IsNil(as, boollib.CallOn(true, okFunc))
	assert.IsNil(as, boollib.CallOrElse(true, okFunc, errFunc))
	assert.IsNotNil(as, boollib.CallOrElse(false, okFunc, errFunc))
}
