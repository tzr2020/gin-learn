// 获取JSON参数

package main

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/json", func(c *gin.Context) {
		// 获取请求主体的字节切片数据
		b, _ := c.GetRawData()

		// 反序列化
		var m map[string]interface{}
		json.Unmarshal(b, &m)

		c.JSON(200, m)
	})

	r.Run()
}
