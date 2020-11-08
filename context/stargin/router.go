package stargin

import (
	"fmt"
	"net/http"
)

type Router struct {
	handlers map [string]Handler
}

func newRouter() *Router {
	return &Router{handlers: make(map[string]Handler)}
}

func (r *Router)handle(c *Context)  {
	key :=c.Request.Method + ">>"+c.Request.URL.Path
	if handler,ok :=r.handlers[key];ok{
		handler(c)
	}else{
		c.Writer.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(c.Writer,"404 NOT FOUND:%s\n",c.Request.URL)
	}
}

func (r *Router)addRoute(method string,path string,handler Handler)  {
	key :=method + ">>"+ path
	r.handlers[key]=handler
}

