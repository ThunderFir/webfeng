package feng

import (
	"net/http"
)

type HandleFunc func(c *Context)

type Engine struct {
	router *router
}

func (engine *Engine) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	c := newContext(writer, request)
	engine.router.handle(c)

	// engine 入口原型 收束请求匹配对应的handler
	//key := request.Method + "-" + request.URL.Path
	//if handler, ok := engine.router[key]; ok {
	//	handler(writer, request)
	//} else {
	//	fmt.Fprintf(writer, "404")
	//}
}

func New() *Engine {
	return &Engine{router: newRouter()}
}

func (engine *Engine) addRoute(method string, pattern string, handler HandleFunc) {
	engine.router.addRoute(method, pattern, handler)
}

func (engine *Engine) Get(pattern string, handler HandleFunc) {
	engine.addRoute("GET", pattern, handler)
}

func (engine *Engine) Post(pattern string, handler HandleFunc) {
	engine.addRoute("POST", pattern, handler)
}

func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}
