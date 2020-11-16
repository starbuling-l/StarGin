package stargin

import "log"

/**
	实现错误处理机制 避免 panic 直接宕机
*/

func Recovery()Handlerfunc  {
	return func(c *Context) {
		defer func() {
			if err:=recover(); err != nil {
				log.Printf("%s\n",err)
			}
		}()
		c.Next()
	}
}