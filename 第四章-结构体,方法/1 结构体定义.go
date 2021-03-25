package main

import "fmt"

/* 结构体是一些列属性组成的复合数据类型，每个属性都具有名称
语法：
	type TypeName formatter
*/

// 基于现有数据类型定义结构体
type Counters int

var num int = 10

// 不同类型之间不能进行运算， 但是可以进行转换后使用
// 错误写法 c2 := num + Counters
// 正确写法 c2 := Counters(num) + Counters

type User map[string]string // 基于现有类型定义

type Counter = int // 别名类型定义

type Callback func() error // 函数类型定义

func main() {
	var counter Counter
	fmt.Printf("%T,%v\n", counter, counter)
	// var user User
	// user["id"] = "1" //无法直接赋值，需要使用make初始化之后才能使用
	var user User = make(User)
	user["id"] = "1"
	fmt.Printf("%T,%v\n", user, user)

	// 函数类型使用
	callbacks := map[string]Callback{}
	callbacks["users"] = func() error {
		fmt.Println("add")
		return nil
	}
	// 调用 函数类型结构体
	callbacks["users"]()

	// 特殊类型 别名
	var (
		r rune
		b byte
	)
	fmt.Printf("%T, %T", r, b)
}
