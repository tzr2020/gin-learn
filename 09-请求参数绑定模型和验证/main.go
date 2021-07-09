// 请求参数绑定模型和验证

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type Student struct {
	Name string `form:"name" uri:"name" json:"name" binding:"required"`
	Age  int8   `form:"age" uri:"age" json:"age" binding:"required,ageVali"`
	// required：内置验证器，请求参数非空
	// ageVali：自定义验证器，年龄参数不小于18
}

// 自定义请求参数验证器
// 年龄参数验证器
func ageVali(fl validator.FieldLevel) bool {
	// 通过反射机制获取请求参数
	if fl.Field().Interface().(int8) < 18 {
		return false // 不通过验证
	}
	return true // 通过验证
}

func main() {
	r := gin.Default()

	// 注册自定义请求参数验证器
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("ageVali", ageVali)
	}

	// 表单参数绑定模型
	// name=zhangsan&age=11
	r.POST("/student", func(c *gin.Context) {
		var s Student
		if err := c.ShouldBind(&s); err != nil {
			c.JSON(200, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"data": s,
		})
	})

	// 查询字符串参数绑定模型
	// /student?name=zhangsan&age=11
	r.GET("/student", func(c *gin.Context) {
		var s Student
		if err := c.ShouldBindQuery(&s); err != nil {
			c.JSON(200, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"data": s,
		})
	})

	// 路径参数绑定模型
	// /student/zhangsan/11
	r.GET("/student/:name/:age", func(c *gin.Context) {
		var s Student
		if err := c.ShouldBindUri(&s); err != nil {
			c.JSON(200, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"data": s,
		})
	})

	// JSON参数绑定模型
	// {
	// 	"name":"zhangsan",
	// 	"age":11
	// }
	r.POST("/student/json", func(c *gin.Context) {
		var s Student
		if err := c.ShouldBindJSON(&s); err != nil {
			c.JSON(200, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"data": s,
		})
	})

	r.Run()
}
