package main

import (
	"fmt"
	"strings"
)

func fors(action string) string {
	// 是否包含字符，左边是对象，右边是匹配的字符
	fmt.Println(strings.Contains(action, "a"))
	// 字符切割, 右边是匹配的字符
	ret := strings.Split(action, "s")
	for i, v := range ret {
		fmt.Println(i, v)
	}
	return "ok"
}

func whiles() {
	fmt.Println("123")
}

func main() {
	var action string
	fmt.Println("请输入您的字符")
	fmt.Scan(&action)
	fors(action)
}
