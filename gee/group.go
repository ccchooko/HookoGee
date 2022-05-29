package gee

import "log"

type RouterGroup struct {
	*router
	prefix      string
	middlewares []HandlerFunc // support middleware
}

// 所有的group公用一个router
func newRootGroup() *RouterGroup {
	return &RouterGroup{
		router: newRouter(),
	}
}

// Engine会继承此方法
func (group *RouterGroup) Group(prefix string) *RouterGroup {
	prefix = group.prefix + prefix
	newGroup := &RouterGroup{
		prefix: prefix,
		router: group.router,
	}
	return newGroup
}

// Engine会继承此方法
func (group *RouterGroup) addRoute(method, comp string, handler HandlerFunc) {
	pattern := group.prefix + comp
	log.Printf("Route %4s - %s", method, pattern)
	group.router.addRouter(method, pattern, handler)
}

// Engine会继承此方法
func (group *RouterGroup) GET(pattern string, handler HandlerFunc) {
	group.addRoute("GET", pattern, handler)
}

// Engine会继承此方法
func (group *RouterGroup) POST(pattern string, handler HandlerFunc) {
	group.addRoute("POST", pattern, handler)
}

