package main

import (
	"fmt"
	"time"
)

type Users struct { //值类型
	id         int
	name       string
	addr       string
	tel        string
	birthday   time.Time
	created_at time.Time
	updated_at time.Time
	flag       bool
	status     int
	deleted_at *time.Time
}

func main() {
	// 指针
	var t *time.Time
	fmt.Println(t) // 零值nil

	// 定义User类型的变量
	var u Users // 所有属性的零值组成的一个Users类型变量值
	fmt.Printf("%T\n", u)
	fmt.Printf("%v\n", u) // 拿到结构体初始化
	fmt.Printf("%#v\n", u)

	// 赋值 按照顺序 需要给所有属性都进行赋值
	u = Users{1, "alinx", "贵州", "188xxxxxx", time.Now(), time.Now(), time.Now(), false, 0, nil}
	fmt.Printf("%#v\n", u)

	// 指定名称进行赋值
	u = Users{id: 2, name: "alinx"}
	fmt.Printf("%#v\n", u)

	// 属性访问
	fmt.Println(u.id)
	fmt.Println(u.name)
	fmt.Println(u.birthday.Clock())

	// 值类型 修改
	u.id = 999999
	fmt.Println(u.id)

	// 指针类型
	var u3 *Users
	fmt.Printf("%T, %v\n", u3, u3)
	// 指针赋值
	u3 = &u
	fmt.Println("%#v\n", u3)

	// new 申请空间并对元素使用0值进行初始化，取地址，赋值
	var u4 *Users = new(Users)
	var u5 *Users = &Users{} // 类似new
	fmt.Printf("%#v\n", u4)
	fmt.Printf("%#v\n", u5) // 初始化值取0值

}
