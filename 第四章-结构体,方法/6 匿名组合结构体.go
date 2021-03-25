package main

import (
	"fmt"
	"time"
)

type Addr struct {
	province string
	street   string
	no       string
}

type Tel struct {
	prefix string
	number string
}

// 匿名组合结构体
type User struct {
	id       int
	name     string
	Addr     // 只指定类型，会自动定义属性名为类型名
	Tel      // 只指定类型
	birthday time.Time
}

func main() {
	// 初始化使用
	var user User = User{Addr: Addr{province: "guizhou"}}
	fmt.Printf("%T,%#v\n", user, user)
	fmt.Println(user.province)
	user.province = "广东省"
	fmt.Println(user.province)

	// 访问和修改
	user.Addr.province = "贵州省"
	fmt.Println(user.Addr.province)

	fmt.Println(user)

}
