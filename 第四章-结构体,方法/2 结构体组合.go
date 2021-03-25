package main

import (
	"fmt"
	"time"
)

// 组合类型结构体
type User struct {
	// 定义结构体属性
	id       int
	name     string
	addr     string
	tel      string
	birthday time.Time
}

func main() {
	// 定义结构体类型的变量
	var user User
	// 零值是由各元素的零值组成的一个结构体的变量
	fmt.Printf("%T,%#v\n", user, user)
	// 字面量初始化方式( 必须严格按照结构体定义顺序指定 ，且每一个属性都必须指定值)
	user = User{10, "alinx", "xxxx", "138", time.Now()}
	fmt.Printf("%T,%#v\n", user, user)

	// 按照属性名定义字面量,可以无序
	user = User{id: 1, name: "alinxsss", addr: "999", tel: "112123123"}
	fmt.Printf("%T, %#v\n", user, user)

	// 属性访问和修改
	// 访问
	fmt.Println(user.id)
	fmt.Println(user.name)
	fmt.Println(user.tel)
	// 修改
	user.id = 9999
	user.name = "good"
	user.birthday = time.Now()
	fmt.Printf("%T, %#v\n", user, user)

}
