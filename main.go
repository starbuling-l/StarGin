package main

/*
(1) /
curl -i http://localhost:9999/
HTTP/1.1 200 OK
Date: Sun, 01 Sep 2019 08:12:23 GMT
Content-Length: 19
Content-Type: text/html; charset=utf-8
<h1>hello stargin</h1>

(2) v1
$ curl -i http://localhost:9999/v1/
HTTP/1.1 200 OK
Date: Mon, 12 Aug 2019 18:11:07 GMT
Content-Length: 18
Content-Type: text/html; charset=utf-8
<h1>Hello StarGin</h1>

(3)
$ curl "http://localhost:9999/v1/hello?name=star"
hello star, you're at /v1/hello

(4)
$ curl "http://localhost:9999/v2/hello/star"
hello star, you're at /hello/star

(5)
$ curl "http://localhost:9999/v2/login" -X POST -d 'username=star&password=123'
{"password":"123","username":"star"}

(6)
$ curl "http://localhost:9999/xxx"
404 NOT FOUND: /xxx
*/

import (
	"github.com/starbuling-l/star-web/stargin"
	"log"
	"net/http"
)

func main() {
	test := stargin.New()
	test.GET("/", func(c *stargin.Context) {
		c.Html(http.StatusOK, "<h1> hello stargin </h1>")
	})

	test.POST("/login", func(c *stargin.Context) {
		c.Json(http.StatusOK, stargin.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	test.GET("/hello", func(c *stargin.Context) {
		log.Print("c.path = ",c.Path)
		c.String(http.StatusOK, "hello %s you are at %s\n", c.Query("name"), c.Path)
	})

	test.GET("/hello/:name", func(c *stargin.Context) {
		c.String(http.StatusOK, "hello %s you are at %s\n", c.Param("name"), c.Path)
	})
	test.GET("static/*filename", func(c *stargin.Context) {
		c.Json(http.StatusOK, stargin.H{
			"filename": c.Param("filename"),
		})
	})

	v1 := test.Group("/v1")
	{
		v1.GET("/", func(c *stargin.Context) {
			c.Html(http.StatusOK, "<h1>Hello StarGin</h1>")
		})
		v1.GET("/hello", func(c *stargin.Context) {
			// /hello?name=star
			c.String(http.StatusOK, "hello %s ,you'are at %s\n", c.Query("name"), c.Path)
		})
	}

	v2 := test.Group("/v2")
	{
		v2.GET("/hello/:name", func(c *stargin.Context) {
			c.String(http.StatusOK, "hello %s ,you'are at %s\n", c.Param("name"), c.Path)
		})

		v2.POST("/login", func(c *stargin.Context) {
			log.Print("username = ",c.PostForm("username"))
			c.Json(http.StatusOK, stargin.H{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})

		})
	}
	test.Run(":9999")
}
