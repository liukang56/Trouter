package middleware

import (
	"Trouter"
	"log"
)

func Logger() Trouter.Middleware {
	return func(next Trouter.ContextFunc) Trouter.ContextFunc {
		return func(c *Trouter.Context) {
			log.Println(c.Request.URL)
			log.Println(c.Request.Header)
			log.Println(c.Request.Method)
			log.Println(c.Request.RemoteAddr)
			next(c)
		}
	}
}
