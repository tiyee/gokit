package swt_test

import (
	"github.com/tiyee/gokit/components/swt"
	"github.com/tiyee/gokit/internal/assert"
	"testing"
)

func TestSwt(t *testing.T) {
	is := assert.NewAssert(t, "TestSwt")
	basic := &swt.BasicPayload{
		UserId: 121,
		Status: 1,
		Name:   "abc",
	}
	encryptStr, err := swt.New(basic)
	assert.IsNil(is, err)
	ba := swt.BasicPayload{}
	err = swt.Parse(encryptStr, &ba)
	assert.IsNil(is, err)
	assert.IsNotNil(is, ba)
	assert.Equal(is, basic.UserId, ba.UserId)
	assert.Equal(is, basic.Status, ba.Status)
	assert.Equal(is, basic.Name, ba.Name)

}
