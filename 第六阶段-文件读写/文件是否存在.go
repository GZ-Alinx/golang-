package main

import (
	"fmt"
	"os"
)

func ma(path string) bool {
	_, err := os.Lstat(path)
	return !os.IsNotExist(err)
}

func main() {
	pa := os.Args[1]
	if ma(pa) {
		fmt.Println("文件存在")
	} else {
		fmt.Println("文件不存在")
	}
}
