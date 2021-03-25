package main

import (
	"fmt"
	"time"
)

func main() {
	c := func() {
		fmt.Println("匿名函数")
	}
	fmt.Printf("%T\n", c)

	// 匿名结构体, 一般用于配置类
	var user struct {
		id       int
		name     string
		tel      string
		addr     string
		birthday time.Time
	}
	fmt.Printf("%T,%#v\n", user, user)

	user.id = 123
	user.name = "sadfasdf"
	fmt.Printf("%T,%#v\n", user, user)

	// 初始化，零值，字面量  使用
	user = struct {
		id       int
		name     string
		tel      string
		addr     string
		birthday time.Time
	}{111, "alinx", "123123", "sdddd", time.Now()} // 结构体类型
	fmt.Printf("%T,%#v\n", user, user)

	// 定义时 就进行初始化

	// var u2 struct {
	// 	id   int
	// 	name string
	// } = struct {
	// 	id   int
	// 	name string
	// }{123, "alinx"}
	// fmt.Printf("%T,%#v\n", u2, u2)

	// 简写 指针加&即可 ，  一般作为数据传递
	var u2 = struct {
		id   int
		name string
	}{id: 111, name: "ddd"}
	fmt.Printf("%T,%#v\n", u2, u2)
	// var u2 = &struct {
	// 	id   int
	// 	name string
	// }{id: 111, name: "ddd"}
	// fmt.Printf("%T,%#v\n", u2, u2)
}
