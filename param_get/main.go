package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//GET传参，格式：/test?first_name=li&last_name=cat
	r.GET("/test", func(c *gin.Context) {
		firstName := c.Query("first_name") // 是 c.Request.URL.Query().Get("first_name") 的简写
		lastName := c.DefaultQuery("last_name", "default_name")
		c.String(http.StatusOK, "%s,%s", firstName, lastName)
	})
	r.Run(":8080")
}
