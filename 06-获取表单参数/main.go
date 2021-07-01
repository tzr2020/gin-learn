// 获取表单参数

package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.POST("/user/search", func(c *gin.Context) {
		// 获取表单参数
		username := c.PostForm("username")

		// username := c.DefaultPostForm("username", "default")

		// username, ok := c.GetPostForm("username")
		// if !ok {
		// 	username = "default"
		// }

		c.JSON(200, gin.H{
			"username": username,
		})
	})

	r.Run()
}
