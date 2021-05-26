package main

import "fmt"

func main() {
	 // 定义切片

	// 定义加赋值
	var slice1 = []string{"123","456","789"}
	fmt.Printf("%T, %#v", slice1, slice1)
	s := slice1[0:1]
	fmt.Printf("%T, %#v\n", s,s)
	fmt.Printf("%T, %#v\n", slice1, slice1)
	// 定义变量为切片类型
	var name []string
	fmt.Printf("%T, %#v", name,name)

	// 数组被切片后 依然保持指针的方式读取 外部操操作会影响数据的数据
	var array1 = [...]string{"1","2","3","4"}
	fmt.Printf("%T, %#v", array1, array1)

	slice2 := array1[1:2]
	slice2[0] = "23123"
	fmt.Printf("%T, %#v\n", slice2, slice2)
	fmt.Printf("%#v\n", array1)
	slice2 = append(slice2, "23124")
	slice2 = append(slice2, "23123")
	slice2 = append(slice2, "23122")
	slice2 = append(slice2, "23121")
	fmt.Printf("\n长度：%T, %#v, %d\n %d\n", slice2, slice2, len(slice2), len(array1))




	// 动态创建切片
	Goslice := make([]string,10,10)
	fmt.Printf("%T, %#v, %d\n", Goslice, Goslice, len(Goslice))

	// 切片的append
	Goslice = append(Goslice, "9999999 ") // 默认将值添加到最后的位
	fmt.Printf("%T, %#v, %d\n", Goslice, Goslice, len(Goslice))
	Goslice = append(Goslice, "0000000 ")
	fmt.Printf("%T, %#v, %d\n", Goslice, Goslice, len(Goslice))


	// 遍历切片
	for index,value := range Goslice {
		fmt.Printf("key:%s value: %s",index,value)
	}
}