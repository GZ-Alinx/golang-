package main

import "fmt"

// 函数在调用是会在堆栈中开辟内存空间，进行存放数据，函数调用结束后会释放
// 变量的生命周期
// 作用域： 代码使用的范围
// 闭包
func addBase(base int) func(int) int {
	return func(num int) int {
		return base + num
	}
}

func gos() {
	add2 := addBase(2) // 闭包调用 1
	fmt.Printf("%T\n", add2)
	fmt.Println(add2(10)) // 闭包调用 2
}

// 作用域
func main() {
	gos() // 闭包调用

	// 以下作用域
	name := "alinx"        //变量
	nums := []int{1, 2, 3} // 切片
	func(pname string, pnums []int) {
		fmt.Println(pname, pnums) // 1
		pname = "goood"
		pnums = []int{9, 8}
		fmt.Println(pname, pnums) // 2
	}(name, nums) // 匿名函数自调用
	fmt.Println(name, nums) // 3

	/*
		执行结果
		alinx [1 2 3]  // 1
		goood [1 2 3]  // 2
		alinx [1 2 3]    // 3
	*/

	// 概念
	// 值传递 的时候都是值拷贝 引用传递(c/c++)
	// 值类型 => 值拷贝
	// 引用类型 => 值拷贝(指针/地址) => 可能会影响函数外部的变量
}
