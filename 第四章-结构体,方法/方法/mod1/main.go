package main

import (
	"fmt"
	"mod1/mods"
)

type User struct {
	id    int
	name  string
	tel   string
	email string
}

// 定义User的方法 返回用户名
func (u User) GetName() string {
	// u: 调用者的值拷贝
	// 可以使用 dlv 进行单步调试
	return "Getname:" + u.name
}

// 修改名称的方法
func (u User) SetName(name string) {
	u.name = name
}

// 指针接收者方法
func (u *User) PsetName(name string) {
	u.name = name
}

// 3. 方法之内不需要访问、需要改 t对象的属性 t可以去掉
func (User) No() {
	fmt.Println("no")
}

func main() {
	u := User{name: "alinx"}
	fmt.Println(u.GetName()) // 调用方法  对象.方法(参数)
	u.SetName("arun111")
	u.PsetName("arun") // 指针方法调用
	// (&u).PsetName("arun222") // 语法糖 和商议中写法类似
	fmt.Println(u.name)

	// GetName
	p := &User{name: "alinx888"}
	fmt.Println(p.GetName()) // 语法糖 (*p).GetName() 编译过程中自动转换
	p.No()                   // 方法之内不需要访问、需要改 t对象的属性 t可以去掉

	// 外部包方法调用
	mo := mods.GetName("88888888888888")
	fmt.Println(mo)
}
