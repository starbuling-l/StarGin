package main

/*
(1)
$ curl -i http://localhost:9999/
HTTP/1.1 200 OK
Date: Mon, 12 Aug 2019 16:52:52 GMT
Content-Length: 18
Content-Type: text/html; charset=utf-8
<h1> hello stargin </h1>

(2)
$ curl curl "http://localhost:9999/hello?name=star"
hello star you are at /hello

(3)
$ curl "curl "http://localhost:9999/hello/star""
hello star you are at /hello/star

(4)
$ curl "http://localhost:9999/static/css/star.css"
{"filename":"css/star.css"}

(5)
$ curl "http://localhost:9999/xxx"
404 NOT FOUND: /xxx
*/

import (
	"github.com/starbuling-l/star-web/stargin"
	"net/http"
)

func main()  {
	test := stargin.New()
	test.GET("/", func(c *stargin.Context) {
		c.Html(http.StatusOK,"<h1> hello stargin </h1>")
	})

	test.POST("/login", func(c *stargin.Context) {
		c.Json(http.StatusOK, stargin.M{
			"username":c.PostForm("username"),
			"password":c.PostForm("password"),
		})
	})

	test.GET("/hello", func(c *stargin.Context) {
		c.String(http.StatusOK,"hello %s you are at %s\n",c.Query("name"),c.Path)
	})

	test.GET("/hello/:name", func(c *stargin.Context) {
		c.String(http.StatusOK,"hello %s you are at %s\n",c.GetParam("name"),c.Path)
	})
	test.GET("static/*filename", func(c *stargin.Context) {
		c.Json(http.StatusOK,stargin.M{
			"filename":c.GetParam("filename"),
		})
	})
	test.Run(":9999")
}