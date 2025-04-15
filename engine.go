package Trouter

import (
	"net/http"
)

type Middleware func(ContextFunc) ContextFunc

type Engine struct {
	Routefunc   map[string]ContextFunc
	Middlewares []Middleware
}

func NewEngine() *Engine {
	return &Engine{
		Routefunc:   make(map[string]ContextFunc),
		Middlewares: make([]Middleware, 0),
	}
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	url := r.Method + "-" + r.URL.Path
	s := e.Routefunc[url]
	s(&Context{w, r})
}

func (e *Engine) Use(middleware Middleware) {
	e.Middlewares = append(e.Middlewares, middleware)
}

func (e *Engine) AddRouter(method string, pattern string, handler ContextFunc) {

	for _, m := range e.Middlewares {
		handler = m(handler)
	}

	url := method + "-" + pattern
	e.Routefunc[url] = handler
}

func (e *Engine) GET(pattern string, handler ContextFunc) {
	e.AddRouter("GET", pattern, handler)
}

func (e *Engine) POST(pattern string, handler ContextFunc) {
	e.AddRouter("POST", pattern, handler)
}

func (e *Engine) DELETE(pattern string, handler ContextFunc) {
	e.AddRouter("DELETE", pattern, handler)
}

func (e *Engine) PUT(pattern string, handler ContextFunc) {
	e.AddRouter("PUT", pattern, handler)
}

func (e *Engine) Group(prefix string) *Group {
	return &Group{
		prefix:      prefix,
		middlewares: make([]Middleware, 0),
		engine:      e,
		parent:      nil,
	}
}
