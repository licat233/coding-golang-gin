package main

import (
	"github.com/gin-gonic/gin"
)

//Person 结构体
type Person struct {
	Age     uint   `form:"age" binding:"required,gt=10"`
	Name    string `form:"name" binding:"required"`
	Address string `form:"address" binding:"required"`
}

func main() {
	r := gin.Default()
	r.GET("/testing", func(c *gin.Context) {
		var person Person
		if err := c.ShouldBind(&person); err != nil {
			c.String(500, "%v", err)
			return
		}
		c.String(200, "%v", person)
	})
	r.Run()
}
