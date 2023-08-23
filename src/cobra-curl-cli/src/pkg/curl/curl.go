package curl

import (
	"cobra-curl-cli/pkg/define"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"cobra-curl-cli/pkg/db"
)

// Run 运行函数
func Run(tCurl define.TCurl, iNum int) (err error) {

	clineName := fmt.Sprintf("%s_%d", tCurl.ClientName, iNum)
	var lastTime time.Time

	for i := 0; i < tCurl.Times; i++ {
		//fmt.Printf("curl access uri [%s], 第 [%d] 次。\n", tCurl.Uri, i+1)
		data, err := curl(tCurl.Uri, tCurl.TimeOut)
		if err != nil {
			fmt.Printf("curl access uri [%s], 第 [%d] 次。 失败：[%v] \n", tCurl.Uri, i+1, err)
		}

		// 是否需要记录到数据库
		if tCurl.SaveDB {
			DBerr := db.SendDb(data, clineName, lastTime)
			if DBerr != nil {
				fmt.Println("SendDB err, TestTool not return!")
				continue
				//return errors.New("SendDB err")
			}
		} else {
			fmt.Printf("[%s] curl access uri [%s], 第 [%d] 次。 成功 data：[%v] \n",
				clineName, tCurl.Uri, i+1, data)
		}
		time.Sleep(time.Second * time.Duration(tCurl.Intervals))
		if err == nil {
			lastTime = time.Now()
		}
	}
	return
}

func curl(uri string, timeout int) (data []byte, err error) {

	// 创建一个自定义的Transport，并忽略证书验证
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	// 创建一个自定义的 HTTP Client，设置超时时间为 5 秒
	httpClient := &http.Client{
		Timeout:   time.Duration(timeout) * time.Second,
		Transport: tr,
	}

	// 发送 GET 请求
	response, err := httpClient.Get(uri)
	if err != nil {
		fmt.Println("Error while sending GET request:", err)
		return
	}
	defer response.Body.Close()

	// 读取响应内容
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error while reading response:", err)
		return
	}

	// 输出响应内容
	data = body
	//fmt.Println("Response:", string(data))
	return
}
