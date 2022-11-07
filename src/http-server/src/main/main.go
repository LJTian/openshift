package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"src/headFunc"
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
	r.GET("/", headFunc.PrintHead)
	r.GET("/edit-file", headFunc.EditFile)

	r.Run(":" + "8080")

	return
}
