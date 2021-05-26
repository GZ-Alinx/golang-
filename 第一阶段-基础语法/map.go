package main

import "fmt"

func main() {

	// 字典的定义第一种 使用make函数构造好map后再进行数据读写操作
	mymap := make(map[int]string)
	fmt.Printf("%T, %#v\n", mymap,mymap)

	// map 值添加
	//mymap[0] 指定key 并赋值

	mymap[0] = "123"
	mymap[1] = "123"
	mymap[2] = "123"
	mymap[3] = "123"
	mymap[999] = "99999"
	fmt.Printf("%T, %#v\n", mymap,mymap)

	// 获取值 根据key获取对应的值
	fmt.Println(mymap[999])


	// 字典的定义第二种 在声明时初始化数据
	mymap2 := map[string]string{
		"name": "alinx",
		"age": "28",
		"site": "CN",
	}
	fmt.Printf("%T, %#v\n", mymap2, mymap2)


	// 判断值是否存在map中
	mate, ok := mymap2["name"]
	if ok {
		fmt.Print(mate)
	}else{
		fmt.Println("值不存在")
	}

	// 便利都是通过for range完成的

	for k,v := range mymap2 {
		fmt.Printf("key:%s value: %s",k,v)
	}
}
