package main

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"os"
)

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		_, _ = os.Stderr.WriteString("your message here")
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/pushover/webhook", func(c *gin.Context) {
		body, _ := ioutil.ReadAll(c.Request.Body)
		_, _ = os.Stderr.WriteString(string(body))

		c.JSON(200, gin.H{
			"message": string(body),
		})
	})
	_ = r.Run(":80")
}
