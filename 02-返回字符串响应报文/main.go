// 返回字符串响应报文

package main

import "github.com/gin-gonic/gin"

func main() {
	// 创建默认路由引擎
	r := gin.Default()

	// 设置路由
	r.POST("/user/login", doLogin)

	// 开启HTTP服务，默认监听端口8080
	r.Run()
}

// 处理器函数
func doLogin(c *gin.Context) {
	// 获取请求参数
	username := c.PostForm("username")
	password := c.PostForm("password")

	// 返回字符串响应报文
	c.String(200, "username: %s\npasswoed: %s\n", username, password)
}
