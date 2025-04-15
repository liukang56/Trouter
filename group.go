package Trouter

type Group struct {
	prefix      string
	middlewares []Middleware
	engine      *Engine
	parent      *Group
}

func (g *Group) Group(prefix string) *Group {
	return &Group{
		prefix:      prefix,
		parent:      g,
		engine:      g.engine,
		middlewares: g.middlewares,
	}
}

func (g *Group) Use(middleware Middleware) {
	g.middlewares = append(g.middlewares, middleware)
}

func (g *Group) Register(method, pattern string, handler ContextFunc) {
	for _, m := range g.middlewares {
		handler = m(handler)
	}

	group := g
	for group != nil {
		pattern = group.prefix + pattern
		group = group.parent
	}
	g.engine.AddRouter(method, pattern, handler)
}

func (g *Group) GET(pattern string, handler ContextFunc) {
	g.Register("GET", pattern, handler)
}

func (g *Group) POST(pattern string, handler ContextFunc) {
	g.Register("POST", pattern, handler)
}

func (g *Group) DELETE(pattern string, handler ContextFunc) {
	g.Register("DELETE", pattern, handler)
}
func (g *Group) PUT(pattern string, handler ContextFunc) {
	g.Register("PUT", pattern, handler)
}
