package base2


/**
	实现自己的逻辑拦截 http 请求
 */
import (
	"fmt"
	"log"
	"net/http"
)

type Engine struct {}

//重载了 ServerHttp 方法
func (engin *Engine) ServeHTTP(writer http.ResponseWriter,request *http.Request) {
	//panic("implement me")
	switch request.URL.Path {
		case "/":
			fmt.Fprintf(writer,"url.path = %q\n",request.URL.Path)
		case "/hello":
			for k,v :=range request.Header {
				fmt.Fprintf(writer,"header[%q] = %q\n",k,v)
			}
		default:
			fmt.Fprintf(writer,"404 NOT FOUND:%s\n",request.URL)
	}
}

func main()  {
	engine := new(Engine)
	log.Fatal(http.ListenAndServe(":9999",engine))
}