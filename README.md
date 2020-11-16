# StarGin
自实现的一个 golang  web 框架


  测试：
   
    (1) /
    $ curl -i http://localhost:9999/
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
    
    (7)
    $ curl "http://localhost:9999/panic"
    2020/11/16 12:44:54 [0] /panic in 0s
     