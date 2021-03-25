package main

import (
	"fmt"
	"time"
)

// 面向对象：  封装，继承，多态

// 结构体类似与面向对象中的类定义
/*
结构体 =>  类
	构造函数 => 创建类对于的实例

	方法



*/
type User struct {
	id       int
	name     string
	tel      string
	addr     string
	birthday time.Time
}

//  New函数定义  ，定义一个结构体

func NewUser(id int, name, tel, addr string, birthday time.Time) User {
	return User{id, name, tel, addr, birthday}
}

// 指针类使用方式
// func NewUser(id int, name, tel, addr string, birthday time.Time) *User {
// 	return &User{id, name, tel, addr, birthday}
// }
func main() {
	// 结构体指针
	// 指针类型,直接取地址 &引用
	// var user *User = &User{id: 999, name: "alinx"}
	//  *User 可以省略
	var user = &User{id: 888, name: "Love"}
	fmt.Printf("%T\n%#v", user, user)

	// 结构体  new函数初始化
	user = new(User)
	fmt.Printf("%T\n,%#v\n", user, user)

	u2 := new(User)
	fmt.Printf("%T,\n%#v\n", u2, u2)

	// 属性访问
	fmt.Println(u2.name)
	u2.id = 9999999999
	u2.name = "gooooood"
	fmt.Println(u2.name, u2.id)

	u3 := NewUser(1, "2222", "", "", time.Now())
	fmt.Printf("%T,%#v\n", u3, u3)
}
