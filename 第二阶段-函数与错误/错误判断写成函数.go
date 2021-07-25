package main

import (
	"fmt"
	"os"
)

func errret(err error) bool {
	if err != nil {
		return false
	}
	return true
}


func main() {
	file,err := os.OpenFile("a.txt", os.O_WRONLY|os.O_APPEND, 755)
	if !errret(err) {
		fmt.Println("打开异常")
	}
	file.Close()
	fmt.Println("结束函数")
}
