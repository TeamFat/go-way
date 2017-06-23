package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func GetNews() {
	doc, err := goquery.NewDocument("https://gocn.io/")
	if err != nil {
		fmt.Println(err)
	}
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		match := "GoCN每日新闻("
		if strings.Contains(s.Text(), match) {
			url, _ := s.Attr("href")
			fmt.Println(match, url)
			os.Exit(0)
		}
	})
}
func main() {
	GetNews()
}
