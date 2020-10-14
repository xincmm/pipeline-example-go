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

	r.GET("/ping", func(c *gin.Context) {
		log.Print("测试西亚")

		f, err := os.OpenFile("/go/src/main.log",
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}
		defer f.Close()

		logger := log.New(f, "prefix", log.LstdFlags)
		logger.Println("text to append")
		logger.Println("more text to append")

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
