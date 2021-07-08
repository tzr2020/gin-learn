// 请求参数绑定

package main

import (
	"github.com/gin-gonic/gin"
)

type User struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func main() {
	r := gin.Default()

	r.POST("login/form", func(c *gin.Context) {
		var u User

		if err := c.ShouldBind(&u); err != nil {
			c.JSON(200, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.String(200, "username: %v\npassword: %v\n", u.Username, u.Password)
	})

	r.POST("login/json", func(c *gin.Context) {
		var u User

		if err := c.ShouldBindJSON(&u); err != nil {
			c.JSON(200, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.String(200, "username: %v\npassword: %v\n", u.Username, u.Password)
	})

	r.Run()
}
