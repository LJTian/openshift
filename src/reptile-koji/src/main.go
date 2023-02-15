package main

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"strconv"
	"strings"
)

var c *colly.Collector

type UtccpSysSet struct {
	Name    string
	IsExist bool
}

var utccpSet []UtccpSysSet
var sysSet []string
var exist = []string{
	"utccp-clients",
	"utccp-hyperkube",
	"utccp-install",
	"utccp-helm",
	"utccp-odo",
	"utccp-operator-sdk",
	"utccp-opm",
	"utccp-pipelines-client",
	"utccp-serverless-clients",
	"cri-o",
	"cri-tools",
	"openvswitch2.16",
	"openvswitch-selinux-extra-policy",
}
var Max int

func initColly() {
	// Instantiate default collector
	c = colly.NewCollector()

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		// fmt.Println("Visiting", r.URL.String())
	})
}

func ExecFuncOfUtccp(e *colly.HTMLElement) {
	class := e.Attr("href")
	if strings.HasPrefix(class, "packageinfo?") {
		data := UtccpSysSet{
			e.Text, false,
		}
		utccpSet = append(utccpSet, data)
	}
}

func ExecFuncOfPageNum(e *colly.HTMLElement) {
	num, err := strconv.Atoi(e.Text)
	if err != nil {
		return
	}
	if num > Max {
		Max = num
	}
}

func ExecFuncOfSys(e *colly.HTMLElement) {
	class := e.Attr("href")
	if strings.HasPrefix(class, "packageinfo?") {
		sysSet = append(sysSet, e.Text)
	}
}

func main() {

	// 初始化架构默认内容
	initColly()
	// On every a element which has href attribute call callback
	c.OnHTML("a[href]", ExecFuncOfUtccp)
	// Start scraping on https://hackerspaces.org
	// 填写需要爬虫的URL
	err := c.Visit("https://koji.uniontech.com/koji/packages?tagID=1630")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	//fmt.Println(utccpSet)

	c.OnHTML("option[value]", ExecFuncOfPageNum)
	c.Visit("https://koji.uniontech.com/koji/packages?tagID=2442")

	c.OnHTML("a[href]", ExecFuncOfSys)
	// Start scraping on https://hackerspaces.org
	for i := 0; i < Max; i++ {
		URL := fmt.Sprintf("https://koji.uniontech.com/koji/packages?start=%d&tagID=2442&order=package_name&inherited=1", 50*i)
		err = c.Visit(URL)
	}

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("检查结果集")

	// 差集标记
	for k, v := range utccpSet {
		for _, v2 := range sysSet {
			if v.Name == v2 {
				utccpSet[k].IsExist = true
			}
		}
	}
	for k, v := range utccpSet {
		for _, v2 := range exist {
			if v.Name == v2 {
				utccpSet[k].IsExist = true
			}
		}
	}

	// 输出结果
	for _, v := range utccpSet {
		if !v.IsExist {
			fmt.Println(v.Name)
		}
	}
}
