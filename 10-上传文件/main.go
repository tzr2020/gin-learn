// 上传文件

package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/file/upload", func(c *gin.Context) {
		f, err := c.FormFile("file")
		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}

		dst := fmt.Sprintf("./%s", f.Filename)
		c.SaveUploadedFile(f, dst)

		c.JSON(200, gin.H{
			"message": fmt.Sprintf("%s 上传成功", f.Filename),
		})
	})

	r.Run()
}
