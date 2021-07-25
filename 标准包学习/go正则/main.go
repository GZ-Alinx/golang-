package main

import (
	"fmt"
	"regexp"
)

func main() {
	// 正则

	/*
		四个函数
			判断是否匹配
			查找字符串
			调换字符串
			分割字符串


		开头	^
		结尾	$
		[a-z]
		[a-zA-Z0-9]
		\d 0-9
		\w  0-9a-zA-Z_
		\D \d取反 非数字
		\W \w取反 不包含[^0-9a-zA-Z_]

		数量限定符
			0或1个		?
			至少1个		+
			任意多个	*
			最少n个		{n,}
			最多m个		{,m}
			n个到m个	{n,m}

		分组
			()

		特殊字符转义
			\\

		贪婪  	尽最大努力匹配
		非贪婪	只匹配一个


	*/

	// 匹配
	text := "aaaa102101a, 123"
	matched, err := regexp.MatchString(`\d{1,10}`, text)   // 是否包含数字
	matched2, err := regexp.MatchString(`^\d{1,10}`, text) // 是否包含以数字开头的字符
	fmt.Println(matched, err, matched2)

	// 邮件正则
	text2 := "123123@163.com"
	mails, err := regexp.MatchString(`^\w{1,32}@\w*[.]((net)||(com)||(cn))$`, text2)
	// `^\w*@\w*\.\w*$`
	fmt.Println(mails)

	// 查找
	text3 := "我的邮箱是： alinx@163.com, 我的qq邮箱： 1135189@qq.com"
	reg1, err := regexp.Compile(`^\w{1,32}@\w*[.]((net)||(com)||(cn))$`)
	reg2 := regexp.MustCompile(`^\w{1,32}@\w*[.]((net)||(com)||(cn))$`)
	fmt.Println(reg1.FindString(text3))
	fmt.Println(reg2.FindAllString(text3, -1))

	// 替换
	fmt.Println(reg1.ReplaceAllString(text3, "*@*.*"))

	// 分割

}
