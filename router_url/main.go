package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//URL传参，标识符(:)或者(*)
	r.GET("/:name/:id", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"name": c.Param("name"),
			"id":   c.Param("id"),
		})
	})
	r.GET("/:name/:id/*action", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"name":   c.Param("name"),
			"id":     c.Param("id"),
			"action": c.Param("action"),
		})
	})

	r.Run()
}
