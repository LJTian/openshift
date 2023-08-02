package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// 生成随机字符串
func generateRandomString(length int) string {

	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func main() {

	nodeName := os.Getenv("NodeName")
	if nodeName == "" {
		nodeName = "Node-" + generateRandomString(5)
	}

	// 创建一个Gin路由引擎
	router := gin.Default()

	// 禁用控制台颜色，将日志写入文件时不需要控制台颜色。
	gin.DisableConsoleColor()

	// 如果需要同时将日志写入文件和控制台，请使用以下代码。
	gin.DefaultWriter = io.MultiWriter(os.Stdout)

	log.Printf("[GIN-debug] node name is [%s]\n", nodeName)

	// 心跳
	router.GET("/ping", func(c *gin.Context) {
		start := time.Now()

		c.Next() // 执行下一个中间件或请求处理函数

		// 在响应头中添加X-Processing-Time字段，记录请求处理时间
		processingTime := time.Since(start)
		c.Header("X-Processing-Time", processingTime.String())

		// 在响应JSON中添加请求时间和服务器时间
		c.JSON(http.StatusOK, gin.H{
			"message":        "pong",
			"processingTime": processingTime.String(),
			"serverTime":     time.Now().Format(time.RFC3339Nano), // 获取服务器当前时间并格式化为RFC3339格式
			"node name":      nodeName,
			"code":           http.StatusOK,
		})
	})

	// 定义一个GET请求的路由处理函数
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message":   "Hello, Gin!",
			"node name": nodeName,
			"code":      http.StatusOK,
		})
	})

	// 添加健康检查路由
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":    "ok",
			"node name": nodeName,
			"code":      http.StatusOK,
		})
	})

	// 启动HTTP服务器并监听在默认端口(8080)
	router.Run()
}
