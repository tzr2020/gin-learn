// 上传和返回文件

package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/uploadfile", func(c *gin.Context) {
		// 设置内存限制，默认是32MiB
		// r.MaxMultipartMemory = 8 << 20 // 8 MiB

		// 读取上传的文件
		f, err := c.FormFile("file")
		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}
		// fmt.Printf("%#v\n", f)

		// 保存文件到本地
		dst := fmt.Sprintf("./%s", f.Filename) // 拼接文件名
		err = c.SaveUploadedFile(f, dst)
		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}

		// c.JSON(200, gin.H{
		// 	"message": fmt.Sprintf("%s上传成功", f.Filename),
		// })

		// 返回本地的文件
		c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", f.Filename))
		c.File(dst) // 使用Go内置的文件服务
	})

	r.POST("/uploadfiles", func(c *gin.Context) {
		form, err := c.MultipartForm()
		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}
		fs := form.File["files"]
		for _, f := range fs {
			fmt.Println(f.Filename)

			// 保存文件到本地
			// c.SaveUploadedFile(f, dst)
		}
	})

	r.Run()
}
