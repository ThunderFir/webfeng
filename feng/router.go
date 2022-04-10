package feng

import "log"

type router struct {
	handlers map[string]HandleFunc
}

func newRouter() *router {
	return &router{handlers: make(map[string]HandleFunc)}
}

func (r *router) AddRouter(method string, path string, handler HandleFunc) {
	key := method + "-" + path
	log.Printf("Route %4s - %s", method, path)
	r.handlers[key] = handler
}

func (r *router) handle(c *Context) {
	key := c.Req.Method + "-" + c.Path
	if handler, ok := r.handlers[key]; ok {
		handler(c)
	} else {
		c.String(StatusNotFound, "404 not found path: %s", c.Path)
	}
}
