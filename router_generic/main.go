package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//泛绑定，标识符（*）
	r.GET("/user/*action", func(c *gin.Context) {
		c.String(200, "hello world")
	})
	r.Run()
}
