package main

import (
	"fmt"
	"net/http"
	stargin "tar-web/base/base3/stargin"
)

func main()  {
	//初始化引擎
	engine := stargin.New()

	//实现静态路由
	engine.GET("/" , func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer,"url.path = %q/n", request.URL.Path)
	})

	engine.GET("/hello", func(writer http.ResponseWriter, request *http.Request) {
		for k,v:=range request.Header {
			fmt.Fprintf(writer,"header[%q] = %q/n",k,v)
		}
	})

	//启动实例
	engine.Run(":9999")
}
