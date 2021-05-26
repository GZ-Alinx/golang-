package main

import (
	"fmt"
	"strings"
)

func main() {
	// 文本处理

	fmt.Println("strings学习")
	fmt.Println(strings.Split("1 2 3", " "))       // 文本切割, 后面是切割标识
	fmt.Println(strings.Contains("123456", "1"))   // 左边字符中是否包含右边的字符 返回值 bool  包含报告"1"是否在"123"之内
	fmt.Println(strings.ContainsAny("123", "1"))   // ContainsAny报告char中的任何Unicode代码点是否在s之内
	fmt.Println(strings.HasPrefix("Golang", "Go")) // 是否是已 Go 开头的字符 bool返回值
	fmt.Println("是否以某字符开头:",strings.HasPrefix("Golang", "Co"))
	fmt.Println("是否以某字符结尾:", strings.HasSuffix("main.windows项目", "windows项目")) // 判断是否以go 结尾的字符

	fmt.Println(strings.Count("1234123123123123123123123", "1")) // 计数，左边存在多少右边的数值
	fmt.Println(strings.Repeat("重复的字符  ", 3))                    // 左边的字符重复，右边为次数int类型
	// 批量替换Replace  参数1： 源字符  参数2： 需要被替换的字符  参数3： 替换后的字符
	fmt.Println(strings.Replace("oldkey oldkey2 oldkeys5 old ", "old", "LXXX", -1)) // 批量替换, 替换次数小于0 相当于没有限制
	fmt.Println(strings.ReplaceAll("oldkey oldkey2 oldkeys5 old ", "old", "L"))     // 批量全部替换 相当于-1
	fmt.Println(strings.ToLower("GoLand"))                                          // 将字符转换为小写
	fmt.Println(strings.ToUpper("golang python java"))                              // 将字符转换为大写
	s := []string{"123", "alinx", "bbb"}
	fmt.Println(strings.Join(s, "000000")) // 字符串拼接 右边参数是拼接参数
}
