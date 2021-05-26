package main

import (
	"container/list"
	"fmt"
)

func main() {
	// 列表的定义
	var name list.List // 生成列表类型
	names := list.New() // 返回列表对应的指针
	fmt.Printf("%T,%#v\n", name,name)
	fmt.Printf("%T,%#v\n", names,names)

	 for i := 1; i < 10; i++ {
	 	names.PushBack(i)
	 }

	 //插入值
	 first :=  names.PushFront(0)

	 // 删除值
	 names.Remove(first)

	 // 遍历值

	 // Front函数获取列表的头元素 Next() 函数依次往下遍历
	 for l := names.Front(); l != nil ; l = l.Next() {
		fmt.Print(l.Value, " ")
	}
}
