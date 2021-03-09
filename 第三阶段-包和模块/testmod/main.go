package main

// 导入math包进行使用
import (
	"fmt"
	"testmod/execs"
	"testmod/math" // 模块名称/目录路径
)

func main() {
	fmt.Println("hello")
	fmt.Println(math.Add(1, 2)) // 包名称 + 函数名称
	fmt.Println(execs.Start())
}
