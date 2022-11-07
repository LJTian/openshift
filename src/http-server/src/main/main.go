package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func main() {

	var nodeName string
	if os.Getenv("NodeName") != "" {
		nodeName = os.Getenv("NodeName")
	} else {
		nodeName = "NULL"
	}
	fmt.Println("node Name is", nodeName)
	hostName, _ := os.Hostname()
	fmt.Println("HostName is", hostName)

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"nodeName": nodeName,
			"hostName": hostName,
		})
	})
	r.GET("/print_head", HeadFunc)

	r.Run(":" + "8080")

	return
}

// HeadFunc 返回请求报文head内容
func HeadFunc(ctx *gin.Context) {

	resp := make(gin.H, 0)
	for k, v := range ctx.Request.Header {
		//fmt.Println(k, v)
		resp[k] = v
	}
	ctx.JSON(http.StatusOK, resp)
}
