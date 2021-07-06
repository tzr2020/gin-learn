// 返回字符串响应报文

package main

import "github.com/gin-gonic/gin"

func main() {
	// 创建默认路由引擎
	r := gin.Default()

	// 设置路由
	r.GET("/index", index)

	// 开启HTTP服务，默认监听端口8080
	r.Run()
}

// 处理器函数
func index(c *gin.Context) {
	// 返回字符串响应报文
	c.String(200, "%s\n", "Hello, world!")
}
