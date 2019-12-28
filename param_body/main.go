package main

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/test", func(c *gin.Context) {
		//获取body的参数，格式：curl -X POST "http://127.0.0.1:8080/test" -d "first_name=li&last_name=cat"
		bodyByts, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			c.Abort()
		}
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyByts))
		firstName := c.PostForm("first_name")
		lastName := c.DefaultPostForm("last_name", "default_last_name")
		c.String(http.StatusOK, "%s,%s,%s", firstName, lastName, string(bodyByts))
	})
	r.Run()
}
