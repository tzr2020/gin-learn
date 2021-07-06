// 获取查询字符串参数

package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	// /user/search?username=zhangsan
	r.GET("/user/search", func(c *gin.Context) {
		// 获取查询字符串参数
		username := c.Query("username")

		// username := c.DefaultQuery("username", "default")

		// username, ok := c.GetQuery("username")
		// if !ok {
		// 	username = "default"
		// }

		c.JSON(200, gin.H{
			"username": username,
		})
	})

	r.Run()
}
