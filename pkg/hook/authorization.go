package hook

import (
	"github.com/tiyee/gokit/pkg/component"
	"github.com/tiyee/gokit/pkg/consts"
	"github.com/tiyee/gokit/pkg/engine"
)

func Authorization(c *engine.Context, fn func(ctx *engine.Context)) {
	ck := c.Request.Header.Cookie(consts.CookieKey)
	//c.SetCookie("muse", []byte(""), time.Second*5)
	if len(ck) < 32 {
		//c.Redirect("https://muse.tiyee.cn/2/login", 302)
		c.AjaxError("请先登录", 403, 1)
		return
	}
	jwt := &component.JWT{}
	if err := jwt.Decrypt(ck); err == nil {
		c.SetUserValue("jwt", jwt)
		fn(c)
	} else {
		c.AjaxError("请先登录", 403, 2)
	}

}
func Authorize(c *engine.Context) {
	ck := c.Request.Header.Cookie(consts.CookieKey)
	if len(ck) < 32 {
		c.AjaxError("请先登录", 403, 1)
		c.Abort()
		return
	}
	jwt := &component.JWT{}
	if err := jwt.Decrypt(ck); err == nil {
		c.SetUserValue("jwt", jwt)
	} else {
		c.AjaxError("请先登录", 403, 2)
		c.Abort()
	}
}
func AuthorizeAdmin(c *engine.Context) {
	ck := c.Request.Header.Cookie(consts.AdminCookieKey)
	if len(ck) < 32 {
		c.AjaxError("请先登录", 405, 1)
		c.Abort()
		return
	}
	jwt := &component.JWT{}
	if err := jwt.Decrypt(ck); err == nil {
		c.SetUserValue("jwt", jwt)
	} else {
		c.AjaxError("请先登录", 405, 2)
		c.Abort()
	}
}
