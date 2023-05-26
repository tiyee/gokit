package handles

import (
	"github.com/tiyee/gokit/pkg/engine"
	"net/http"
)

func Signature(c *engine.Context) {
	echostr := c.Query("echostr")
	c.String(http.StatusOK, echostr)
}
