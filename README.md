# StarGin
自实现的一个 golang  web 框架

- day 01：

  测试：(1)
     $ curl -i http://localhost:9999/
     
     HTTP/1.1 200 OK
     
     Date: Mon, 12 Aug 2019 16:52:52 GMT
     
     Content-Length: 18
     
     Content-Type: text/html; charset=utf-8
     
     <h1> hello stargin </h1>
     
     (2)
     $ curl "http://localhost:9999/hello?name=star"
     
     hello star you are at /hello
     
     (3)
     $ curl "http://localhost:9999/hello/star"
     
     hello star you are at /hello/star
     
     (4)
     $ curl "http://localhost:9999/static/css/star.css"
     
     {"filename":"css/star.css"}
     
     (5)
     $ curl "http://localhost:9999/xxx"
     
     404 NOT FOUND: /xxx
