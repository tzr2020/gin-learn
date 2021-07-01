// 返回HTML文档响应报文

package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	// 加载模板文件
	// r.LoadHTMLFiles("templates/posts/index.html", "templates/users/index.html")
	r.LoadHTMLGlob("templates/**/*")

	r.GET("/posts/index", func(c *gin.Context) {
		c.HTML(200, "posts/index.html", gin.H{
			"title": "posts/index",
		})
	})

	r.GET("/users/index", func(c *gin.Context) {
		c.HTML(200, "users/index.html", gin.H{
			"title": "users/index",
		})
	})

	r.Run()
}
