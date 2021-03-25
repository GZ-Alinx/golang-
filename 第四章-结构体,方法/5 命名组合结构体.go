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

// 命名组合(嵌入)
type User struct {
	id       int
	name     string
	addr     *Addr // 组合Addr结构体
	tel      *Tel  // 组合Tel结构体
	birthday time.Time
}

func main() {
	// 初始化使用
	var user User = User{id: 1, name: "alinx", addr: &Addr{}, tel: &Tel{}, birthday: time.Now()}
	fmt.Printf("%T,%#v", user, user)
	// 写入 和读取  如果是指针就加指针进行调用即可  &Addr
	var users User = User{addr: &Addr{province: "贵州省"}}

	// 访问
	fmt.Println(users.addr.province)
	// fmt.Printf("%T,%#v", users, users)

}
