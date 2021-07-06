/*
	中间件
	拦截请求的特殊处理器函数
*/

package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.New()

	// 全局使用Gin自带的中间件Logger和Recovery
	r.Use(gin.Logger(), gin.Recovery())

	r.GET("/index", func(c *gin.Context) {
		c.String(200, "%s\n", "Hello, world!")
		panic("这是一个恐慌")
	})

	r.Run()
}
