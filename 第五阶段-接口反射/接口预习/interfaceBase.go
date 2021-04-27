package main

import "fmt"

type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human  // 匿名组合
	school string
	loan   float32
}

type Employee struct {
	Human
	company string
	money   float32
}

// Human 实现Sayhi方法
func (h Human) Sayhi() {
	fmt.Println("Hi 我是%s,我的电话是%s", h.name, h.phone)
}

func (h Human) Sing(lyrics string) {
	fmt.Println("la la la ", lyrics)
}

func (e Employee) Sayhi() {
	fmt.Println("Hi I am %s, 我的工作是%s ,联系我 %s", e.name, e.company, e.phone)
}

type Men interface {
	SayHi()
	Sing(lyrics string)
}

func main() {
	mike := Student{Human{"Mike", 28, "0755-87881211"}, "MIT", 9999}
	alinx := Student{Human{"Alinx", 26, "0766-12312312"}, "CNMIT", 20000}
	lixiong := Student{Human{"lx", 66, "13877777777"}, "GGG", 99999}

	// 定义Men类型的变量
	var i Men
	i = mike
	fmt.Println("This is mike , a Student")
	i.SayHi()
	i.Sing("November rain")

	i = alinx
	fmt.Println("This is Alinx, a Student:")
	i.Sing("alinx")
	i.SayHi()
}
