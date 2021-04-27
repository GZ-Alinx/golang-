package main

import "fmt"

func main() {

	//切片删除数据
	data1 := []string{"golang", "java", "python", "C#", "C++"}
	//data1=append(data1[:2],data1[2:]...)
	for i, v := range data1 {
		fmt.Println(i, v)
		if v == "java" {
			data1 = append(data1[:i], data1[i+1:]...)
		}
	}
	// 删除特定的值
	//data1 = append(data1[:2], data1[2+1:]...)
	fmt.Println(data1)

	//append操作
	c := []int{1, 2, 3, 4, 5}
	//c=append(c[:i],c[i+1:]...)  删除中间一个元素:（c[i]）
	//c=append(c[:i],c[i+N:]...)  删除中间N个元素:（c[i:I+N-1])
	//先取c[:1]之前的数据，在追加c[2:]之后的数据,即删除c[1]
	c = append(c[:1], c[2:]...)
	fmt.Println(c) //[1 3 4 5]
}
