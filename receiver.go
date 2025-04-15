package Trouter

import (
	"encoding/json"
	"net/http"
)

type ContextFunc func(*Context)

type Context struct {
	Writer  http.ResponseWriter
	Request *http.Request
}

func (c *Context) Bind(v any) error {
	return json.NewDecoder(c.Request.Body).Decode(v)
}

func (c *Context) Query(key string) string {
	return c.Request.URL.Query().Get(key)
}
