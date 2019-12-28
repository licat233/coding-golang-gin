package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//POST传参:curl -X POST "http://127.0.0.1:8080/test" -d "first_name=li&last_name=cat"
	r.POST("/test", func(c *gin.Context) {
		firstName := c.PostForm("first_name")
		lastName := c.DefaultPostForm("last_name", "anything")
		c.String(http.StatusOK, "%s,%s", firstName, lastName)
	})
	r.Run()
}
