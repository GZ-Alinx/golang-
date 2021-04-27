package mods

import "fmt"

type User struct {
	id    int
	name  string
	tel   string
	email string
}

// 定义User的方法 返回用户名
func (u User) GetName(name string) string {
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
