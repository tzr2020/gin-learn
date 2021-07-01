// 请求重定向

package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// HTTP重定向
	r.GET("/index", func(c *gin.Context) {
		c.Redirect(301, "https://www.baidu.com")
	})

	// 路由重定向
	r.GET("/home", func(c *gin.Context) {
		c.Request.URL.Path = "/test"
		r.HandleContext(c)
	})

	r.GET("/test", func(c *gin.Context) {
		c.String(200, "Path: %s", c.Request.URL.Path)
	})

	r.Run()
}
