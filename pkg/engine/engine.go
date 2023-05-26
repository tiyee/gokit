package engine

import (
	"fmt"
	"github.com/tiyee/gokit/pkg/component/log"
	"github.com/tiyee/gokit/pkg/consts"
	"net/http"
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

	return &Context{r: nil, w: nil, handlers: make([]HandlerFunc, 0), index: 0}
}
func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	methodS := r.Method
	pathS := r.URL.Path
	fmt.Println(pathS)
	key := methodS + pathS
	context := e.pool.Get().(*Context)
	context.reset(w, r)
	if route, exist := e.routes[key]; exist {

		for _, fn := range route.BeforeHooks {
			context.handlers = append(context.handlers, fn)
		}
		context.handlers = append(context.handlers, route.HandlerFuncs...)
		for _, fn := range route.AfterHooks {
			context.handlers = append(context.handlers, fn)
		}
		defer func() {
			if err := recover(); err != nil {
				log.Error("recover", log.String("error", fmt.Sprintf("%v", err)))
				context.Error("内部错误", 500)
			}
		}()
		defer func() {
			if err := log.Sync(); err != nil {
				fmt.Println(err.Error())
			}
		}()
		context.index = -1
		context.Next()
		e.pool.Put(context)
	} else {
		context.NotFound()
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

	return http.ListenAndServe(consts.ADDR, e)

}
func (e *Engine) GET(path string, fn ...HandlerFunc) {
	e.setRoute(http.MethodGet, path, fn...)
}
func (e *Engine) POST(path string, fn ...HandlerFunc) {
	e.setRoute(http.MethodPost, path, fn...)
}
func (e *Engine) PUT(path string, fn ...HandlerFunc) {
	e.setRoute(http.MethodPut, path, fn...)
}
func (e *Engine) DELETE(path string, fn ...HandlerFunc) {

	e.setRoute(http.MethodDelete, path, fn...)
}
func (e *Engine) OPTIONS(path string, fn ...HandlerFunc) {
	e.setRoute(http.MethodOptions, path, fn...)
}
func (e *Engine) PATCH(path string, fn ...HandlerFunc) {
	e.setRoute(http.MethodPatch, path, fn...)
}

func (e *Engine) dispatch() {
	for _, route := range e.routes {
		for _, hook := range e.hooks {
			if hook.matcher.Match(route.Method, route.Path) {
				if hook.Pos == PosAhead {
					route.BeforeHooks = append(route.BeforeHooks, hook.HandlerFunc)
				}
				if hook.Pos == PosBehind {
					route.AfterHooks = append(route.AfterHooks, hook.HandlerFunc)
				}
			}
		}
	}
}
