// 获取路径参数

package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	// /user/search/zhangsan
	r.GET("/user/search/:username", func(c *gin.Context) {
		// 获取路径参数
		username := c.Param("username")

		c.JSON(200, gin.H{
			"username": username,
		})
	})

	r.Run()
}
