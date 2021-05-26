package main

import "fmt"

type Printer interface {
	Print(interface{})
}

type FuncCaller struct {
}

//实现Printer的Print方法
func (f FuncCaller) Print(p interface{}){
	fmt.Printf("%T, %#v\n", p,p)
}

func main() {
	var ff Printer
	ff = &FuncCaller{}
	ff.Print("ttttt")
	ff.Print(true)
}
