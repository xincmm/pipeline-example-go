package main

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	log.SetOutput(os.Stdout)
	r := gin.Default()

	log.Print(os.Getenv("MODE"))

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/pushover/webhook", func(c *gin.Context) {
		body, _ := ioutil.ReadAll(c.Request.Body)
		log.Print(string(body))
		c.JSON(200, gin.H{
			"message": string(body),
		})
	})
	_ = r.Run(":80")
}
