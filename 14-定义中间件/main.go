/*
	中间件适合处理一些公共的业务逻辑，比如登录认证、权限校验、数据分页、记录日志、耗时统计等
*/
package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	// 全局使用中间件
	// r.Use(logger)

	// 单个路由使用中间件
	r.GET("/index", logger, index)

	// 单个路由组使用中间件
	userGroup := r.Group("/user")
	userGroup.Use(logger)
	{
		userGroup.GET("/index", index)
	}

	r.Run()
}

// 定义中间件
func logger(c *gin.Context) {
	start := time.Now()
	c.Set("start", start)

	// 继续下一个处理器函数进一步处理请求
	c.Next()
	// 中止下一个处理器函数进一步处理请求
	// c.Abort()

	log.Printf("请求路径: %v\n", c.Request.URL.Path)
	log.Printf("请求处理耗时: %v\n", time.Since(start))
	log.Printf("响应状态码: %v\n", c.Writer.Status())
}

func index(c *gin.Context) {
	start := c.MustGet("start").(time.Time)
	c.String(200, "请求开始处理时间: %v\n请求处理耗时: %v\n",
		start, time.Since(start))
}
