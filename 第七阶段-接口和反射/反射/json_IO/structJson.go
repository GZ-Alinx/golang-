package main

import (
	"encoding/json"
	"fmt"
)

// 基础用法


// 需要设置属性与Json字符串中的对应关系
// Id <=> json.pk
// 结构体属性标签： `` 字符串：key:"value1"


// 结构体 首字母要大写 否则外部无法访问
type User struct {
	Id int 					`json:"id"` // 反序列化
	Name string				`json:"name"`
	IsBoby bool 			`json:"isboby"`
	Scores []float64 		`json:"scores"`
	Phone map[string]string `json:"phone"`
}



func main() {

	// 结构体转换为 json字符串
	user := User{1,"asdfasdfasdf   0   	/*  ",true,[]float64{1,2,3},map[string]string{"a":"1","b":"alinx"}}
	b,err := json.Marshal(user)
	fmt.Printf("%#v,%#v",err,string(b))


	fmt.Println("")
	// json字符串，转换为结构体
	jsonB := `
	{
		"Id": 10,
		"Name": "alinx",
		"IsBoby": false,
		"Phone":{"mobile":"123123"}
	}
	`
	var u User
	err = json.Unmarshal([]byte(jsonB), &u)
	fmt.Printf("%#v,%#v", err,u)
}
