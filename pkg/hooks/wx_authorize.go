package hooks

import (
	"github.com/tiyee/gokit/pkg/component/log"
	"github.com/tiyee/gokit/pkg/consts"
	"github.com/tiyee/gokit/pkg/controllers/offiaccount"
	"github.com/tiyee/gokit/pkg/engine"

	"net/http"
)

func WXAuthorize(c *engine.Context) {
	//if !strings.HasPrefix(c.Request.RequestURI, "/wx") {
	//	c.Next()
	//	return
	//}
	signature := c.Query("signature")
	timestamp := c.Query("timestamp")
	nonce := c.Query("nonce")

	ok := offiaccount.CheckSignature(signature, timestamp, nonce, consts.WxToken)
	if !ok {
		log.Error("[微信接入] - 微信公众号接入校验失败!")
		c.String(http.StatusForbidden, "signature err!")
		return
	}
	c.Next()
}
