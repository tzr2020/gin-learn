// 返回JSON文本响应报文

package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.GET("json", func(c *gin.Context) {
		// c.JSON(200, gin.H{
		// 	"useranme": "zhangsan",
		// 	"password": "123456",
		// })

		user := struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}{
			Username: "lisi",
			Password: "123456",
		}

		c.JSON(200, user)
	})

	r.Run()
}
