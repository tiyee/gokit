package hooks

import (
	"github.com/tiyee/gokit/pkg/component"
	"github.com/tiyee/gokit/pkg/consts"
	"github.com/tiyee/gokit/pkg/engine"
)

func Authorize(c *engine.Context) {
	ck := c.Cookie(consts.CookieKey)
	if len(ck) < 32 {
		c.AjaxError(403, "请先登录", 1)
		c.Abort()
		return
	}
	jwt := &component.JWT{}
	if err := jwt.Decrypt([]byte(ck)); err == nil {
		c.SetUserValue("jwt", jwt)
	} else {
		c.AjaxError(403, "请先登录", 2)
		c.Abort()
	}
}
func AuthorizeAdmin(c *engine.Context) {

	ck := c.Cookie(consts.AdminCookieKey)
	if len(ck) < 32 {
		c.AjaxError(405, "请先登录", 1)
		c.Abort()
		return
	}
	jwt := &component.JWT{}
	if err := jwt.Decrypt([]byte(ck)); err == nil {
		c.SetUserValue("jwt", jwt)
	} else {
		c.AjaxError(405, "请先登录", 2)
		c.Abort()
	}
}
