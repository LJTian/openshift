package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	// 这里false 表示 http 服务，非 https
	r := gin.Default()
	r.GET("/login", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"sub":  "base-000001",
			"name": "base-Name-000001",
		})
	})
	r.GET("/login/fail", func(c *gin.Context) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Error message",
		})
	})

	r.Run(":" + "8080")

	return
}
