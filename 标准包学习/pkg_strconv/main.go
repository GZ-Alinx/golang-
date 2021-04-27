package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 数据类型转换包 类型处理

	fmt.Println(strconv.Atoi("123"))                             // 字符串转换为int类型  Atoi等效于ParseInt（s，10，0），
	fmt.Printf("%T，%#v\n", strconv.Itoa(123), strconv.Itoa(123)) // int整型转换为字符串

	// bool类型 转换
	fmt.Println(strconv.FormatBool(0 < 1)) // 根据值类型返回值
	fmt.Println(strconv.FormatBool(1 < 0))
	fmt.Println(strconv.ParseBool("1"))
	fmt.Println(strconv.ParseBool("0"))

	//  将bool类型转换为字符串
	fmt.Println(strconv.FormatBool(false))

	// 将字符串转换为浮点数
	fmt.Println(strconv.ParseFloat("1.23333", 32))

	// 将字符串转换为int类型 参数1 需要转换的字符串 参数2 进制， 参数3 int(? /8/16/32/64)
	fmt.Println(strconv.ParseInt("123333", 10, 64))
	// Atoi 相当于 ParseInt(s, 10, 0)
	fmt.Println(strconv.Atoi("123"))

	// 浮点数转换为字符串
	fmt.Println(strconv.FormatFloat(1.2222, 'e', 5, 64))
	// 格式标记：
	// 'b' (-ddddp±ddd，二进制指数)
	// 'e' (-d.dddde±dd，十进制指数)
	// 'E' (-d.ddddE±dd，十进制指数)
	// 'f' (-ddd.dddd，没有指数)
	// 'g' ('e':大指数，'f':其它情况)
	// 'G' ('E':大指数，'f':其它情况)

	id := []map[string]string{}
	user := map[string]string{
		"id":   "123",
		"user": "999",
		"add":  "sz",
	}
	id = append(id, user)
	fmt.Println(id)
	for _, v := range id {
		fmt.Println(v["id"])
		fmt.Printf("%T,%s", v["id"], v["id"])
	}
}
