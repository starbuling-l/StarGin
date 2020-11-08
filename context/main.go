package main

/*
(1)
curl -i http://localhost:9999/

HTTP/1.1 200 OK
Content-Type: text/html
Date: Sun, 08 Nov 2020 10:16:02 GMT
Content-Length: 24

<h1> hello stargin </h1>%

(2)
curl -i "http://localhost:9999/hello?name=star"

HTTP/1.1 200 OK
Content-Type: text/plain
Date: Sun, 08 Nov 2020 10:19:15 GMT
Content-Length: 29

hello star you are at /hello

(3)
$ curl "http://localhost:9999/login" -X POST -d 'username=geektutu&password=1234'
{"password":"1234","username":"geektutu"}

(4)
$ curl "http://localhost:9999/xxx"
404 NOT FOUND: /xxx
*/


import (
	"github.com/starbuling-l/StarGin/stargin"
	"net/http"
)

func main()  {
	test :=stargin.New()
	test.GET("/", func(c *stargin.Context) {
		c.Html(http.StatusOK,"<h1> hello stargin </h1>")
	})

	test.POST("/login", func(c *stargin.Context) {
		c.Json(http.StatusOK,stargin.M{
			"username":c.PostForm("username"),
			"password":c.PostForm("password"),
		})
	})

	test.GET("/hello", func(c *stargin.Context) {
		c.String(http.StatusOK,"hello %s you are at %s\n",c.Query("name"),c.Path)
	})
	test.Run(":9999")
}