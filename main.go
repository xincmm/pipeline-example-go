package main

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		println("Hello Rancher")
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/pushover/webhook", func(c *gin.Context) {
		body, _ := ioutil.ReadAll(c.Request.Body)
		println(string(body))
		c.JSON(200, gin.H{
			"message": string(body),
		})
	})
	_ = r.Run(":80")
}
