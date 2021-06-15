package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
)



func main() {
	// id
	// class
	document, _ := goquery.NewDocument("https://www.baidu.com")
	document.Find("table").Each(func(int, section *goquery.Selection) {
		fmt.Println(index)
		fmt.Println(section.Attr("border"))
	})

}
