package handle

import (
	"github.com/tiyee/gokit/pkg/component"
	"github.com/tiyee/gokit/pkg/engine"
	"github.com/tiyee/gokit/pkg/repository/class_repo"
)

func Test(c *engine.Context) {
	component.Logger.Info("test")
	component.Logger.Error("error")
	//if time.Now().Unix() > 100000 {
	//	test(c)
	//	return
	//}
	class, err := class_repo.Class(4313325028647451)
	if err != nil {
		c.AjaxError("error", 111, err.Error())
		return
	}
	m := map[string]string{
		"test": "tes4",
		"x":    class.Name,
	}
	c.AjaxSuccess("ok", m)

}
