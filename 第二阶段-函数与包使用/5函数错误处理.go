package main

import (
	"errors"
	"fmt"
	"strconv"
)

// error 接口类型
// 错误类型定义 第一种
func div1(left, right int) (int, error) {
	if right == 0 {
		return 0, fmt.Errorf("right num is zero") // 初始化error类型变量
	}
	return left / right, nil
}

// 错误类型定义 第二种
func div2(left, right int) (int, error) {
	if right == 0 {
		return 0, errors.New("right num is zero") // 初始化error类型变量
	}
	return left / right, nil
}

func main() {
	// Atoi 转换为int类型
	v, err := strconv.Atoi("123")
	// 调用者直接进行错误处理  err 错误如果为空这为正常 不为空则有错误
	if err == nil {
		fmt.Println(v)
	} else {
		fmt.Printf("转换错误: %s\n", err)
	}
	// 执行结果
	// 123 <nil>

	// 错误分类
	/*
		1. 运行时错误
			可恢复错误 => 重试、忽略
			不可恢复错误 => 程序退出

		2. 语法错误
			改变语法格式

	*/

	fmt.Println(div1(1, 0)) // 调用出发错误  第一种
	fmt.Println(div1(2, 1)) // 调用得到正确 nil
	// 多种写法
	v, err = div1(1, 0)
	if err == nil {
		fmt.Println(v)
	}
	if v, err = div1(1, 0); err == nil {
		fmt.Println(v)
	} else {
		fmt.Println("value is error", v)
	}
	fmt.Println(div2(1, 0)) // 调用出发错误  第二种
	fmt.Println(div2(2, 1)) // 调用得到正确 nil
}
