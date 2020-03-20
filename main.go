package main

import (
	"fmt"
	// "os"
	"strings"

	// "github.com/davecgh/go-spew/spew"
	"github.com/gocolly/colly/v2"
)

func main() {
	// crawl("https://shop366791302.taobao.com/")
	// crawl("https://shop110714426.taobao.com/")
	crawl("https://m2mcoffee.taobao.com/search.htm")
}

func crawl(link string) {
	goods := map[string]bool{}

	// Instantiate default collector
	c := colly.NewCollector(
	// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
	// colly.AllowedDomains("hackerspaces.org", "wiki.hackerspaces.org"),
	)

	// On every a element which has href attribute call callback
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		l := e.Attr("href")
		// Print link
		// Visit link found on page
		// Only those links are visited which are in AllowedDomains
		// if strings.Contains(e.Text, "咖啡") && !strings.Contains(link, l) {
		if strings.Contains(e.Text, "咖啡") && !strings.Contains(link, l) {
			fmt.Printf("Link found: %q -> %s\n", e.Text, l)
			c.Visit(e.Request.AbsoluteURL(l))
			// if !strings.Contains(e.Text, "\n") && !strings.Contains(e.Text, "挂耳") &&
			// 	!strings.Contains(e.Text, "分类") &&
			// 	!strings.Contains(e.Text, "器具") &&
			// 	// !strings.Contains(e.Text, "意式") &&
			// 	!strings.Contains(e.Text, "杯") &&
			// 	!strings.Contains(e.Text, "服务") {
			// 	// len(e.Text) > 30 {
			// 	goods[e.Text] = true
			// }
			// if !strings.Contains(e.Text, "\n") && !strings.Contains(e.Text, "挂耳") &&
			// 	!strings.Contains(e.Text, "分类") &&
			// 	!strings.Contains(e.Text, "器具") &&
			// 	// !strings.Contains(e.Text, "意式") &&
			// 	!strings.Contains(e.Text, "杯") &&
			// 	!strings.Contains(e.Text, "服务") {
			// 	// len(e.Text) > 30 {
			goods[e.Text] = true
			// }
		}
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		// fmt.Println("Visiting", r.URL.String())
	})

	// Start scraping on https://hackerspaces.org
	// c.Visit(link + "/search.htm")
	c.Visit(link)

	// spew.Dump(goods)

	for k, _ := range goods {
		k = strings.Replace(k, "黑咖啡", "", -1)
		k = strings.Replace(k, "咖啡豆", "", -1)
		k = strings.Replace(k, "咖啡", "", -1)
		k = strings.Replace(k, "m2m", "", -1)
		k = strings.Replace(k, "M2M", "", -1)
		k = strings.Replace(k, "手冲", "", -1)
		k = strings.Replace(k, "单品", "", -1)
		k = strings.Replace(k, "新鲜", "", -1)
		k = strings.Replace(k, "coffee", "", -1)
		k = strings.Replace(k, "COFFEE", "", -1)
		k = strings.Replace(k, "新鲜", "", -1)
		k = strings.Replace(k, "烘焙", "", -1)
		k = strings.Replace(k, "精品", "", -1)
		k = strings.Replace(k, " ", "", -1)
		k = strings.Replace(k, "2019", "", -1)
		k = strings.Replace(k, "产季", "", -1)
		k = strings.Replace(k, "KW", "", -1)
		k = strings.Replace(k, "kw", "", -1)
		k = strings.Replace(k, "无糖", "", -1)
		k = strings.Replace(k, "正品", "", -1)
		k = strings.Replace(k, "滴滤", "", -1)
		k = strings.Replace(k, "100g", "", -1)
		k = strings.Replace(k, "227g", "", -1)
		k = strings.Replace(k, "105g", "", -1)
		k = strings.Replace(k, "105", "", -1)
		k = strings.Replace(k, "120g", "", -1)
		k = strings.Replace(k, "270g", "", -1)
		fmt.Println(k)
	}
}
