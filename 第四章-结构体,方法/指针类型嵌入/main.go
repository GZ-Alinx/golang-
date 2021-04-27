package main

import "fmt"

type address struct {
	Province string
	City     string
	Street   string
}

type User struct {
	ID       int
	Name     string
	Tel      string
	*address // 匿名嵌入 只能嵌入一次 多了会重复
}

func main() {
	// 1. 定义变量
	var u User
	fmt.Printf("%#v\n", u)

	// 2. 赋值
	u = User{ID: 666, address: &address{"云南省", "昆明市", "中医药大学"}}
	fmt.Printf("%#v\n", u)
	fmt.Printf("%#v\n", u.address)

	// 3. 属性访问和修改
	fmt.Printf("%#v\n", u.address)
	fmt.Printf("%#v\n", u.City)
	u.address = &address{"广东省", "深圳市", "深圳大学"}
	fmt.Printf("%#v\n", u.address)
	fmt.Printf("%#v\n", u.City)
	fmt.Printf("%#v\n", u.Province)
	fmt.Printf("%#v\n", u.address.Province)
	fmt.Printf("%#v\n", u.address.Province)
}
