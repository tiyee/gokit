package engine

import (
	"regexp"
	"strings"
)

type HookPosition int

const PosAhead HookPosition = 1
const PosBehind HookPosition = 2

type IMatcher interface {
	Match(method, path string) bool
}
type Hook struct {
	matcher     IMatcher
	Pos         HookPosition
	HandlerFunc HandlerFunc
}

type Contain struct {
	sub string
}

func Contains(sub string) *Contain {
	return &Contain{sub: sub}
}

type Reg struct {
	regexp *regexp.Regexp
}

func RegExp(expr string) *Reg {
	reg, err := regexp.Compile(expr)
	if err != nil {
		panic(err.Error())
	}
	return &Reg{regexp: reg}
}
func (r *Reg) Match(method, path string) bool {
	return r.regexp.MatchString(path)

}

type Identically struct {
	path string
}

func Identical(path string) *Identically {
	return &Identically{path: path}
}

func (i *Identically) Match(method, path string) bool {
	return path == i.path
}
func (c *Contain) Match(method, path string) bool {
	return strings.Contains(path, c.sub)
}

type Pre struct {
	prefix    string
	exclusive []string
}

func Prefix(prefix string, exclusive []string) *Pre {

	return &Pre{prefix: prefix, exclusive: exclusive}
}
func (p *Pre) Match(method, path string) bool {
	for _, s := range p.exclusive {
		if s == path {
			return false
		}
	}
	return strings.HasPrefix(path, p.prefix)
}
func (e *Engine) Use(pos HookPosition, matcher IMatcher, fn HandlerFunc) {
	e.hooks = append(e.hooks, Hook{
		matcher:     matcher,
		Pos:         pos,
		HandlerFunc: fn,
	})
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
