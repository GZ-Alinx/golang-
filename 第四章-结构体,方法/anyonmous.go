package main

import "fmt"

func main() {
	// 匿名结构体 不定义名称  不能再定义对应的变量  只定义该种结构的一个变量

	var user struct {
		id   int
		name string
		addr string
		tel  string
	}
	fmt.Printf("%T\n", user)
	fmt.Printf("%3v\n", user)

	user.id = 10
	fmt.Printf("%d\n", user.id)

	// 赋值

	// 按照顺序赋值
	user = struct {
		id   int
		name string
		addr string
		tel  string
	}{1, "alinx", "gz", "188xxx"}
	// 按照key值赋值
	user = struct {
		id   int
		name string
		addr string
		tel  string
	}{1, "alinx", "gz", "188xxx"}
	fmt.Println(user)

	//  以上简写, 常用写法
	var users = struct {
		id   int
		name string
		addr string
		tel  string
	}{1, "alinx", "gz", "188xxx"}
	fmt.Println(users)

	// 指针类型
	u5 := &struct {
		id   int
		name string
	}{1, "alinx"}
	fmt.Println(u5)

	// 后续使用
	// template -- 数据(匿名结构体)  -- 传递给模板

	var users1 = []struct {
		id   int
		name string
	}{
		{1, "kk"},
		{2, "alinx"},
	}
	fmt.Printf("%T\n", users1)
	fmt.Printf("%#v\n", users1)
	// 类型相同 可以简写
	type USERS struct {
		id              int
		name, tel, site string
	}
}
