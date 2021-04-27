package main

import "fmt"

type Address struct {
	Province string
	City     string
	Street   string
}

type User struct {
	ID      int
	Name    string
	Tel     string
	Address // 匿名嵌入 只能嵌入一次 多了会重复
	City    string
}

// 匿名嵌入2个对象 两个对象中有相同的属性名
type Company struct {
	Province string
}

type User2 struct {
	ID      int
	Name    string
	Tel     string
	Address // 匿名嵌入 只能嵌入一次 多了会重复
	City    string
	Company
}

func main() {
	// 1. 定义变量
	var u User
	fmt.Printf("%#v\n", u)

	// 2. 赋值
	u = User{ID: 666, Address: Address{"云南省", "昆明市", "中医药大学"}}
	fmt.Printf("%#v\n", u)

	// 3. 属性访问和修改
	fmt.Println(u.Address.City)
	fmt.Println(u.Address.Province)
	u.Address.City = "曲靖市{}"
	u.City = "曲靖市"
	fmt.Printf("%#v\n", u)
	fmt.Printf("%#v\n", u.City)         // 可以直接访问
	fmt.Printf("%#v\n", u.Address.City) // 可以直接访问

	// 多个匿名嵌入对象 具有相同属性名称，将无法找到，并抛出异常，强制指定使用那个报错
	var u3 User2
	fmt.Printf("%#v\n", u3)
	u3.Company.Province = "深圳地区"
	u3.Address.Province = "深圳地区 good"
	fmt.Printf("%T\n", u3.Company.Province)  // 使用绝对路径
	fmt.Printf("%#v\n", u3.Address.Province) // 使用绝对路径

}
