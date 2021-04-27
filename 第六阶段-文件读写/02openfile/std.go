package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	data := make([]byte, 3)
	n, err := os.Stdin.Read(data)
	fmt.Println(n, err, data)

	os.Stdout.WriteString("你属如的内容为:" + string(data[:n]))
	fmt.Fprintf(os.Stdout, "你输入的内容为:"+string((data[:n])))

	path := "p.txt"
	file, _ := os.Open(path)
	if err != nil {
		log.Fatal(file)
	}
	defer file.Close()

	fmt.Println(file.Readdir(-1))
	//fmt.Println(file.IsDir())
	fmt.Println(os.Create("xxx.txt"))

}
