package stargin

import (
	"log"
	"time"
)

/**
	中间件打印日志
 */
func Logger()Handlerfunc  {
	return func(c *Context) {
		t := time.Now()
		c.Next()
		log.Printf("[%d] %s in %v",c.StatusCode,c.Request.URL,time.Since(t))
	}
}
