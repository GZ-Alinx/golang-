package main

import "fmt"

// 地址结构体
type Address struct {
	Province string
	City     string
	Street   string
}

// 定义用户信息
type User struct {
	ID    int
	Name  string
	Tel   string
	Email string
	Addr  Address  // 值类型 命名嵌入  嵌入Address结构体 即定义Address结构体属性，组合Address结构体
	PAddr *Address // 指针类型 嵌入Address结构体
} // 命名嵌入

func main() {
	// 定义
	var u User
	var addr Address = Address{Province: "贵州", City: "黔西南", Street: "兴义市"}
	u = User{ID: 1, Name: "alinx", Tel: "188xxxxxxxx", Email: "aaaa@qq.com", Addr: addr, PAddr: &Address{"广东省", "深圳市", "福田区"}}

	// 属性修改和访问
	fmt.Println(u)
	fmt.Println(u.ID)
	fmt.Println(u.Addr)
	u.ID = 9999999999
	u.Addr = Address{"广东省", "深圳市", "宝安区"}
	fmt.Println(u.Addr)
	fmt.Printf("%#v\n", u.PAddr) // 指针类型
	// 修改指针类型
	u.PAddr.Street = "宝安区"
	fmt.Printf("%#v\n", u.PAddr) // 指针类型

}
