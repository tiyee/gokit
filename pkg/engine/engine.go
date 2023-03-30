package engine

import (
	"fmt"
	"github.com/tiyee/gokit/pkg/component"
	"github.com/valyala/fasthttp"
	"go.uber.org/zap"
	"strings"
	"sync"
)

type HandlerFunc func(ctx *Context)
type Route struct {
	Method       string
	Path         string
	BeforeHooks  []HandlerFunc
	AfterHooks   []HandlerFunc
	HandlerFuncs []HandlerFunc
}
type Routes map[string]*Route
type Engine struct {
	routes Routes
	addr   string
	pool   sync.Pool
	hooks  []Hook
}

func New() *Engine {
	engine := &Engine{
		routes: make(map[string]*Route, 0),
		addr:   ":3003",
		hooks:  make([]Hook, 0),
	}
	engine.pool.New = func() any {
		return engine.allocateContext()
	}
	return engine

}
func (e *Engine) allocateContext() *Context {

	return &Context{RequestCtx: nil, handlers: make([]HandlerFunc, 0), index: 0}
}
func (e *Engine) HandleFastHTTP(ctx *fasthttp.RequestCtx) {
	methodS := string(ctx.Method())
	pathS := string(ctx.Path())
	key := methodS + pathS

	if route, exist := e.routes[key]; exist {
		context := e.pool.Get().(*Context)
		context.reset(ctx)
		for _, fn := range route.BeforeHooks {
			context.handlers = append(context.handlers, fn)
		}
		context.handlers = append(context.handlers, route.HandlerFuncs...)
		for _, fn := range route.AfterHooks {
			context.handlers = append(context.handlers, fn)
		}
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
				component.Logger.Error("recover", zap.String("error", fmt.Sprintf("%v", err)))
				ctx.Error("内部错误", 500)
			}
		}()
		defer func() {
			if err := component.Logger.Sync(); err != nil {
				fmt.Println(err.Error())
			}
		}()
		context.index = -1
		context.Next()
		e.pool.Put(context)
	} else {
		ctx.NotFound()
	}
}

func (e *Engine) setRoute(method string, path string, fn ...HandlerFunc) {

	path = "/" + strings.Trim(path, "/")
	key := method + path
	e.routes[key] = &Route{
		Method:       method,
		Path:         path,
		HandlerFuncs: fn,
		BeforeHooks:  make([]HandlerFunc, 0),
		AfterHooks:   make([]HandlerFunc, 0),
	}
}
func (e *Engine) Run() (err error) {
	e.dispatch()
	return fasthttp.ListenAndServe(e.addr, e.HandleFastHTTP)
}
func (e *Engine) GET(path string, fn ...HandlerFunc) {
	e.setRoute("GET", path, fn...)
}
func (e *Engine) POST(path string, fn ...HandlerFunc) {
	e.setRoute("POST", path, fn...)
}
func (e *Engine) PUT(path string, fn ...HandlerFunc) {
	e.setRoute("PUT", path, fn...)
}
func (e *Engine) DELETE(path string, fn ...HandlerFunc) {
	e.setRoute("Delete", path, fn...)
}
