package main

import "fmt"

func iflen(action string) string {
	if action == "yes" {
		fmt.Println("your install docker")
		return "install"
	} else if action == "no" {
		fmt.Println("not install docker")
		return "not install"
	} else {
		fmt.Println("选择错误")
		return "error"
	}
}

func main() {
	var action string
	fmt.Print("请选择操作:")
	fmt.Scan(&action)
	iflen(action)
}
