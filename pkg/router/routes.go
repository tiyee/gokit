package router

import (
	"github.com/tiyee/gokit/pkg/engine"
	"github.com/tiyee/gokit/pkg/handles"
	"github.com/tiyee/gokit/pkg/handles/uploader"
	"github.com/tiyee/gokit/pkg/hooks"
)

type IRouter interface {
	POST(path string, fn ...engine.HandlerFunc)
	GET(path string, fn ...engine.HandlerFunc)
	PUT(path string, fn ...engine.HandlerFunc)
	DELETE(path string, fn ...engine.HandlerFunc)
	OPTIONS(path string, fn ...engine.HandlerFunc)
	Use(pos engine.HookPos, matcher engine.IMatcher, fn engine.HandlerFunc)
}

func LoadRouter(r IRouter) {

	r.GET("/wx", handles.Signature)
	r.POST("/wx", handles.Portal)

	r.POST("/uploader/init", uploader.Init)
	r.POST("/uploader/upload", uploader.Upload)
	r.POST("/uploader/merge", uploader.Merge)

	r.Use(engine.PosAhead, engine.Prefix("/wx", []string{}), hooks.WXAuthorize)
}
