package main

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

func main() {
	// Instantiate default collector
	c := colly.NewCollector()

	// On every a element which has href attribute call callback
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		var err error
		link := e.Attr("href")
		// fmt.Println(e.Text)
		// Print link
		if len(e.Text) > 1 {
			format := e.Text[len(e.Text)-1:]
			// fmt.Printf("[%s]\n", format)

			if format == "/" {
				// 目录
				//fmt.Printf("dir Link found: %q -> %s\n", e.Text, link)
				err = c.Visit(e.Request.AbsoluteURL(link))
			} else if e.Text == " Parent Directory" || e.Text == "Parent Directory" {
				// 上层目录
				// fmt.Printf("back Link found: %q -> %s\n", e.Text, link)
				// Visit link found on page
				// Only those links are visited which are in AllowedDomains
				// err = c.Visit(e.Request.AbsoluteURL(link))
			} else {
				fmt.Printf("src Link found: %q -> %s\n", e.Text, e.Request.AbsoluteURL(link))
			}
		} else {
			fmt.Printf("Link found: %q -> %s\n", e.Text, link)
			err = c.Visit(e.Request.AbsoluteURL(link))

		}
		if err != nil && err.Error() != "URL already visited" {
			fmt.Println(err.Error())
			return
		}
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		// fmt.Println("Visiting", r.URL.String())
	})

	// Start scraping on https://hackerspaces.org
	// 填写需要爬虫的URL
	err := c.Visit("https://ftp.redhat.com/pub/redhat/linux/enterprise/8Base/en/RHMT/container-sources/")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
