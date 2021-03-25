package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("main")
	args := os.Args[1:]

	for k, v := range args {
		fmt.Println(k, v)
	}

	if len(args) < 2 {
		fmt.Println("参数不合法")
	}

	if len(args) >= 3 {
		fmt.Println("参数太多了")
	}
}
