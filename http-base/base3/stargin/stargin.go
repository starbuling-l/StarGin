package stargin

import (
	"fmt"
	"net/http"
)

type Handler func(http.ResponseWriter, *http.Request)

type Engine struct {
	router map[string]Handler
}

//实现 ServeHTTP 可以实现自逻辑 如果发现该路径有注册则走注册 handler 否则抛出 404异常
func (engine *Engine) ServeHTTP(writer http.ResponseWriter,request *http.Request) {
	//panic("implement me")
	key := request.Method + ">>" +request.URL.Path
	if Handler,ok:=engine.router[key];ok{
		Handler(writer,request)
	}else{
		writer.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(writer,"404 NOT FOUND:%s\n",request.URL)
	}
}

//初始化方法
func New() *Engine {
	return &Engine{router: make(map[string]Handler)}
}

//注册路由方法
func (engine *Engine)addRoute(method string, pattern string,handler Handler)  {
	key := method + ">>"+ pattern
	engine.router[key] = handler
}

//引擎启动 包装的 ListenAndServe
func (engine *Engine) Run (addr string)(err error)  {
	//实现了接口方法的 struct 可强转为接口类型
	//handler := (http.Handler)(engine) // 手动转换为借口类型
	//log.Fatal(http.ListenAndServe(":9999", handler))
	return http.ListenAndServe(addr,engine) //第二个参数类型为接口类型 http.Handler 要实现其方法

}

// get 方法实现
func (engine *Engine) GET (pattern string ,handler Handler)  {
	engine.addRoute("GET",pattern,handler)
}

//post 方法实现
func (engine *Engine) POST (pattern string,handler Handler)  {
	engine.addRoute("POST",pattern,handler)
}