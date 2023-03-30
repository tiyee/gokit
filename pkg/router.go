package pkg

import (
	"github.com/tiyee/gokit/pkg/engine"
	"github.com/tiyee/gokit/pkg/handle"
	"github.com/tiyee/gokit/pkg/hook"
)

type IRouter interface {
	POST(path string, fn ...engine.HandlerFunc)
	GET(path string, fn ...engine.HandlerFunc)
	PUT(path string, fn ...engine.HandlerFunc)
	DELETE(path string, fn ...engine.HandlerFunc)
	Use(pos engine.HookPosition, matcher engine.IMatcher, fn engine.HandlerFunc)
}

func LoadRouter(r IRouter) {

	r.GET("/test", handle.Test)

	r.Use(engine.PosAhead, engine.Prefix("/teacher", []string{"/teacher/login"}), hook.AuthorizeAdmin)
}
