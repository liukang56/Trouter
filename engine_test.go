package Trouter

import (
	"fmt"
	"net/http"
	"testing"
)

func TestNewEngine(t *testing.T) {
	e := NewEngine()
	g := e.Group("/api")
	g.GET("/sdhj", dshjh)
	g.GET("/dfsd", dshjh)

	//用户组

	http.ListenAndServe(":8080", e)
}

type User struct {
	Name string
}

func dshjh(c *Context) {
	id := c.Query("id")
	//var u User
	//if err := c.Bind(&u); err != nil {
	//	panic(err)
	//}
	fmt.Println(id)
}
