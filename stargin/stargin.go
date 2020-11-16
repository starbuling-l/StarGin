package stargin

import (
	"net/http"
)

type Handlerfunc func(c *Context)

type Engine struct {
	router *Router        //继承路由
	*RouterGroup   //继承分组路由 不加别名 group
	groups []*RouterGroup //分组嵌套
}

//初始化方法
func New() *Engine {
	engine := &Engine{router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

//默认使用 Logger,Recovery 中间件
func Default()*Engine {
	engine:=New()
	engine.Use(Logger(),Recovery())
	return engine
}

//引擎启动 包装的 ListenAndServe
func (engine *Engine) Run(addr string) (err error) {
	//实现了接口方法的 struct 可强转为接口类型
	//handler := (http.Handler)(engine) // 手动转换为借口类型
	//log.Fatal(http.ListenAndServe(":9999", handler))
	return http.ListenAndServe(addr, engine) //第二个参数类型为接口类型 http.Handler 要实现其方法
}

//实现 ServeHTTP 可以实现自逻辑 如果发现该路径有注册则走注册 handler 否则抛出 404异常 分组添加中间件
func (engine *Engine) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	//panic("implement me")
	c := newContext(writer, request)
	for _, group := range engine.groups {
		c.handlers=append(c.handlers,group.middlewares...)
	}
	engine.router.handle(c)
}
