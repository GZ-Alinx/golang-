package main

import (
	"fmt"
	"visib/models"
)

func main() {
	// 访问测试
	fmt.Println(models.Public{})
	public := models.Public{}
	fmt.Println(public.Public)
	fmt.Println(public.Private) //  无法调用
	// fmt.Println(models.privateStruct{})

	pv := models.NewPrivateStruct()
	fmt.Println(pv)
}
