package main

import "fmt"

// go语言中所有的类型都能拥有方法

type vars int

func (f vars) tests() {
	fmt.Println("tests")
}


func main() {
	var a vars
	a.tests()
}
