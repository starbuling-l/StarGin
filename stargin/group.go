package stargin

import "log"

/**
实现路由分组控制(Route Group Control)

例如：
以/post开头的路由匿名可访问。
以/admin开头的路由需要鉴权。
以/api开头的路由是 RESTful 接口，可以对接第三方平台，需要三方平台鉴权。
*/

type RouterGroup struct {
	prefix      string        //前缀
	middlewares []Handlerfunc //todo 分组中间件
	engine      *Engine       //循环嵌套可以实现 分组 和 不分组 两种模式 同时所有实例共用一个 engine
}

func (group *RouterGroup) Group(prefix string) *RouterGroup {
	engine := group.engine
	newGroup := &RouterGroup{
		prefix: group.prefix + prefix,
		engine: group.engine,
	}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

// group 添加中间件
func (group *RouterGroup) Use(middlewares ...Handlerfunc) {
	group.middlewares = append(group.middlewares, middlewares...)
}

//注册路由方法
func (group *RouterGroup) addRoute(method string, comp string, handler Handlerfunc) {
	pattern := group.prefix + comp
	log.Printf("Route %4s - %s", method, pattern)
	group.engine.router.addRoute(method, pattern, handler)
}

// get 方法实现
func (group *RouterGroup) GET(pattern string, handler Handlerfunc) {
	group.addRoute("GET", pattern, handler)
}

//post 方法实现
func (group *RouterGroup) POST(pattern string, handler Handlerfunc) {
	group.addRoute("POST", pattern, handler)
}
