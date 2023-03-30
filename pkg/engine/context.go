package engine

import (
	"encoding/json"
	"errors"
	"github.com/tiyee/gokit/pkg/component"
	"github.com/tiyee/gokit/pkg/consts"
	"github.com/tiyee/gokit/pkg/helps"
	"github.com/valyala/fasthttp"
	"math"
	"time"
)

const abortIndex int8 = math.MaxInt8 >> 1

type JsonRet struct {
	Error   int         `json:"error"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
type JsonRetWithPagination struct {
	Error   int         `json:"error"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Pages   int64       `json:"pages"`
}
type Context struct {
	*fasthttp.RequestCtx
	handlers []HandlerFunc
	index    int8
}

func (c *Context) reset(ctx *fasthttp.RequestCtx) {
	c.RequestCtx = ctx
	c.handlers = c.handlers[:0]
	c.index = 0
}
func (c *Context) Next() {
	c.index++
	for c.index < int8(len(c.handlers)) {
		c.handlers[c.index](c)
		c.index++
	}
}

// IsAborted returns true if the current context was aborted.
func (c *Context) IsAborted() bool {
	return c.index >= abortIndex
}

// Abort prevents pending handlers from being called. Note that this will not stop the current handler.
// Let's say you have an authorization middleware that validates that the current request is authorized.
// If the authorization fails (ex: the password does not match), call Abort to ensure the remaining handlers
// for this request are not called.
func (c *Context) Abort() {
	c.index = abortIndex
}
func (c *Context) GetPostArgsIntZero(key string) int64 {

	return c.GetPostArgsInt(key, 0)
}
func (c *Context) GetQueryArgsIntZero(key string) int64 {

	return c.GetQueryArgsInt(key, 0)
}
func (c *Context) GetPostArgsInt(key string, defaultValue int64) int64 {
	bs := c.PostArgs().Peek(key)
	return helps.BytesToInt64(bs, defaultValue)
}
func (c *Context) GetQueryArgsInt(key string, defaultValue int64) int64 {
	bs := c.QueryArgs().Peek(key)
	return helps.BytesToInt64(bs, defaultValue)
}
func (c *Context) Uid() int64 {
	if jwt, err := c.JWT(); err == nil {
		return jwt.Uid
	} else {
		return 0
	}
}
func (c *Context) JWT() (*component.JWT, error) {
	i := c.UserValue("jwt")
	if i == nil {
		return nil, errors.New("jwt not found")
	}
	if jwt, ok := i.(*component.JWT); ok {
		return jwt, nil
	} else {
		return nil, errors.New("jwt interrupt error")
	}
}
func (c *Context) AjaxError(msg string, errorCode int, data interface{}) {
	ret := JsonRet{
		Message: msg,
		Error:   errorCode,
		Data:    data,
	}
	if data, err := json.Marshal(ret); err == nil {
		c.Success("application/json", data)
	} else {
		str := `{"message":"error","error":1}`
		c.Success("application/json", []byte(str))
	}
}
func (c *Context) AjaxSuccessWithPagination(msg string, pages int64, data interface{}) {
	ret := JsonRetWithPagination{
		Message: msg,
		Error:   0,
		Data:    data,
		Pages:   pages,
	}
	if data, err := json.Marshal(ret); err == nil {
		c.Success("application/json", data)
	} else {
		str := `{"message":"error","error":1}`
		c.Success("application/json", []byte(str))
	}
}
func (c *Context) AjaxSuccess(msg string, data interface{}) {
	ret := JsonRet{
		Message: msg,
		Error:   0,
		Data:    data,
	}
	if data, err := json.Marshal(ret); err == nil {
		c.Success("application/json", data)
	} else {
		str := `{"message":"format err","error":1}`
		c.Success("application/json", []byte(str))
	}
}
func (c *Context) SetCookie(key string, value []byte, expired time.Duration) {
	cookie := &fasthttp.Cookie{}
	cookie.SetDomain(consts.Domain)
	cookie.SetHTTPOnly(true)
	cookie.SetPath("/")
	cookie.SetKey(key)
	cookie.SetExpire(time.Now().Add(expired))
	cookie.SetSecure(true)

	cookie.SetValueBytes(value)
	c.Response.Header.SetCookie(cookie)
}
