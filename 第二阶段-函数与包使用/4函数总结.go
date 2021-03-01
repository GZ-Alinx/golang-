package main

import "fmt"

// 函数是代码块的封装,代码复用的最基本单元

/*
	函数定义
	形参，实参
	func 函数名(参数1 类型，参数2 类型，参数n 类型) (返回值1 类型,返回值2 类型,返回值n 类型) {
		函数体
	}

	// 函数参数类型，返回值类型可以为函数、等其他数据类型


	函数使用
	// 函数可以赋值给变量，可以在数组，切片，映射，结构体中

	匿名函数作为参数到函数中使用
	func start(func() string) func(int) string {
		return func(int) string {
			return "匿名函数参数调用"
		} 
	}
	

	函数-值类型&引用类型
		值类型：数值、布尔、字符串、指针、数组、结构体等
		引用类型： 切片、映射、接口等

		针对值类型可以借助指针修改原值
*/




func test(stop bool, st="123" string) {
	fmt.Println("A")
	if stop {
		fmt.Println("B")
		// 如果没有定义返回值
		/* return 两个作用
		1. 返回值，
		2. 退出循环体
		*/
		return
	}
	fmt.Println("C")
}

func main() {
	test(true)
	fmt.Println("------")
	test(false)
}
