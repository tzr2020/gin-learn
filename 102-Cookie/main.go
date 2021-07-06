package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.GET("/cookie", func(c *gin.Context) {
		// 设置cookie
		c.SetCookie("username", "zhangsan", 3600, "/", "localhost", false, true)

		// 获取cookie
		value, err := c.Cookie("username")
		if err != nil {
			value = "not found"
		}

		c.String(200, "%v\n", value)
	})

	r.Run()
}
