package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	// 获取宏参数
	filePath := os.Getenv("FILE_NAME")
	if filePath == "" {
		filePath = "nfs-mount-ok"
	}
	fmt.Println("filePath is : " + filePath)

	// 创建文件
	file, err := os.Create("/mnt/" + filePath)
	if err != nil {
		fmt.Println("无法创建文件:", err)
		return
	}
	defer file.Close()

	// 写入内容
	content := []byte("Hello, World!")
	_, err = file.Write(content)
	if err != nil {
		fmt.Println("无法写入文件:", err)
		return
	}

	fmt.Println("文件创建并写入成功")

	for {
		time.Sleep(time.Second)
	}
}
