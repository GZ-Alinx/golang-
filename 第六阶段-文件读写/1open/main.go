package main

import (
	"fmt"
	"os"
)

func main() {

	path := "a.txt"
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(file)
}
