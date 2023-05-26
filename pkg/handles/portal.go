package handles

import (
	"fmt"
	"github.com/tiyee/gokit/pkg/engine"
	"io"
	"log"
	"net/http"
)

func Portal(c *engine.Context) {
	data, err := io.ReadAll(c.Request().Body)
	if err != nil {
		log.Println(err.Error())
		return
	}
	fmt.Printf("data: %T %v\n", data, string(data))

	c.String(http.StatusOK, "")
}
