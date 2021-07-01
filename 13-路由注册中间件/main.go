/*
	路由注册中间件
	中间件适合处理一些公共的业务逻辑，比如登录认证、权限校验、数据分页、记录日志、耗时统计等
*/

package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// 定义中间件
func StatCost() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Set("message", "Hello, world!")

		// 下一个处理器函数进一步处理请求
		c.Next()

		cost := time.Since(start)
		log.Println(cost)
	}
}

func main() {
	r := gin.New()

	// 全局路由注册中间件
	r.Use(StatCost())

	r.GET("/test", func(c *gin.Context) {
		message := c.MustGet("message").(string)
		c.JSON(200, gin.H{
			"message": message,
		})
	})

	// 单个路由注册中间件
	r.GET("/test2", StatCost(), func(c *gin.Context) {
		message := c.MustGet("message").(string)
		c.JSON(200, gin.H{
			"message": message,
		})
	})

	// 路由组注册中间件
	// userGroup := r.Group("/user", StatCost())
	// {
	// 	userGroup.GET("index", func(c *gin.Context) {})
	// }

	userGroup := r.Group("/user")
	userGroup.Use(StatCost())
	{
		userGroup.GET("index", func(c *gin.Context) {})
	}

	r.Run()
}
