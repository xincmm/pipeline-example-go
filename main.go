package main

import (
	"flag"
	_ "flag"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"k8s.io/klog/v2"
)

func main() {
	r := gin.Default()
	klog.InitFlags(nil)
	_ = flag.Set("logtostderr", "false")
	_ = flag.Set("log_file", "myfile.log")
	flag.Parse()

	r.GET("/ping", func(c *gin.Context) {
		klog.Info("Hello Rancher")
		klog.Flush()
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/pushover/webhook", func(c *gin.Context) {
		body, _ := ioutil.ReadAll(c.Request.Body)
		klog.Info(string(body))
		klog.Flush()

		c.JSON(200, gin.H{
			"message": string(body),
		})
	})
	_ = r.Run(":80")
}
