package headFunc

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"time"
)

// HeadFunc 返回请求报文head内容
func PrintHead(ctx *gin.Context) {

	resp := make(gin.H, 0)
	for k, v := range ctx.Request.Header {
		//fmt.Println(k, v)
		resp[k] = v
	}
	ctx.JSON(http.StatusOK, resp)
}

// HeadFuncEditFile 返回请求报文head内容
func EditFile(ctx *gin.Context) {

	var fileName string

	if os.Getenv("fileName") != "" {
		fileName = os.Getenv("NodeName")
	} else {
		fileName = "syslog.log"
	}
	fileName = "/var/log/myweb/" + fileName
	fmt.Println("fileName is: ", fileName)

	// If the file doesn't exist, create it, or append to the file
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	hour, minute, second := time.Now().Clock()
	time := fmt.Sprintf("%02d:%02d:%02d\n", hour, minute, second)
	if _, err := f.Write([]byte(time)); err != nil {
		f.Close() // ignore error; Write error takes precedence
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"fileName": fileName,
		"data":     time,
		"status":   "ok",
	})
}

// HeadFunc 返回请求报文head内容
func LogPrint(c *gin.Context) {

	// 获取表单参数
	fmt.Println("loginfo { ")
	formParams := c.Request.PostForm
	for key, values := range formParams {
		for _, value := range values {
			fmt.Printf("%s: %v\n", key, value)
		}
	}

	// 获取JSON参数
	var jsonParams map[string]interface{}
	if err := c.ShouldBindJSON(&jsonParams); err == nil {
		for key, value := range jsonParams {
			fmt.Printf("%s: %v\n", key, value)
		}
	}
	fmt.Println("} ")

	// 返回200响应
	c.Status(200)
}
