package main

import (
	"encoding/json"
	"fmt"
)

// 首先设定json格式 并按照json的字段编写结构体

type Info struct {
	Age int 	`json:"age"`
	Site string `json:"site"`
}

type User struct {
	Name Info `json:"name"`
	Sex bool `json:"sex"`
}


// 序列化  将数据转换为json格式进行处理
func JsonMarshal() []byte{
	user := &User{}
	user.Name.Site = "alinx"
	user.Name.Age = 28
	user.Sex = true

	data, _ :=  json.Marshal(user)
	fmt.Println(string(data))
	return data
}



// 反序列化 将json字符串转换为可调用对象

func JsonUnMarshal() {
	jsondata := JsonMarshal()
	data := &User{}
	Udata := json.Unmarshal([]byte(jsondata), data)
	fmt.Println(Udata)
}





func main() {
	// 掉用  格式和级别一定要规划好 才能更好的写出好的程序
	JsonMarshal()
	JsonUnMarshal()


}