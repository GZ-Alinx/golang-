package main

import (
	"fmt"
)

// 用户管理，使用切片的方式进行存储用户信息
var User []string = []string{}

func add(user string) bool {
	User = append(User, user)
	return true
}

func del() {
	User = append(User)
}

func modify() {

}

func Query() {
	User = append(User)
}

func main() {

	fmt.Println("欢迎使用用户管理系统")
	fmt.Printf("%T", User)
}
