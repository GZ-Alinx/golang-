package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	// ioutil
	// io

	data, _ := ioutil.ReadFile("rdwr.go")
	fmt.Println(string(data))

	files, _ := ioutil.ReadDir(".")
	for _, info := range files {
		fmt.Println(info.Name())
	}
	ioutil.WriteFile("ioutil.txt", []byte("xxx"), os.ModePerm)

	f1, _ := os.Create("f1.txt")
	f2, _ := os.Create("f2.txt")
	m := io.MultiWriter(f1, f2)
	// 同时写入两份文件
	m.Write([]byte("gopoooo"))

	f1, _ = os.Open("go.mod")
	f2, _ = os.Open("mnain.go")
	r := io.MultiReader(f1, f2)
	io.Copy(os.Stdout, r)

}
