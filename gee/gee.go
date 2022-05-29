package gee

import (
	"net/http"
)

type HandlerFunc func(c *Context)

type Engine struct {
	*RouterGroup
}

func New() *Engine {
	//return &Engine{router: newRouter()}
	group := newRootGroup()
	engine := &Engine{
		RouterGroup: group,
	}
	return engine
}

//func (engine *Engine) Group(prefix string) *RouterGroup {
//	group := engine.RouterGroup.Group(prefix)
//	engine.groups = append(engine.groups, group)
//	return group
//}

//func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
//	engine.router.addRouter(method, pattern, handler)
//}
//
//func (engine *Engine) GET(pattern string, handler HandlerFunc) {
//	engine.addRoute("GET", pattern, handler)
//}
//
//func (engine *Engine) POST(pattern string, handler HandlerFunc) {
//	engine.addRoute("POST", pattern, handler)
//}

func (engine *Engine) Run(addr string) (err error) {
	// ListenAndServe(addr string, handler Handler)
	// engine也就是第二个参数，handler，必须是实现ServerHTTP(http.ResponseWriter, *http.Request)
	// 也就是说，只要传入任何实现了 ServerHTTP 接口的实例，所有的HTTP请求，就都交给了该实例处理了。
	return http.ListenAndServe(addr, engine)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := newContext(w, r)
	engine.router.haddle(c)
}
