// 设置路由与路由组

package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	// 普通路由
	r.GET("/index", func(c *gin.Context) {})
	r.POST("/login", func(c *gin.Context) {})
	r.Any("/test", func(c *gin.Context) {})

	r.NoRoute(func(c *gin.Context) {})

	// 路由组
	userGroup := r.Group("/user")
	{
		userGroup.GET("/index", func(c *gin.Context) {})
		userGroup.GET("/login", func(c *gin.Context) {})
		userGroup.POST("/login", func(c *gin.Context) {})
	}

	r.Run()
}
