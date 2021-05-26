package main



/*
json 格式化的字符串
数值： 	number(int/float) => 1.1 10 -10
字符串： string => ""
布尔：	bool => true/false
数组：	[]Type/[Length] =>
字典：	map[Ktype]Value => {"key1":value}

Go数据类型 <=> Gob二进制字符串
Go数据类型 <=> Json字符串
Go数据类型 <=> xml字符串
Go数据类型 <=> Protobuffer


持久化/网络传输
=> 序列化 => Go数据类型 => 字符串(文本/二进制)
=> 反序列化 => 字符串(文本/二进制) => Go数据类型

编码规则 => 库(包) 类似json库就是一种

查询
go doc encoding/json
*/

import (
	"encoding/json"
	"fmt"
)
func main() {
	/*
	Marshal 序列化
	Unimarshal 反序列化
	 */
	var (
		number float64 = 1.1
		name string = "alinx"
		isbody bool = true
		user []string = []string{"alinx","golang","china"}
		scores map[string]int = map[string]int{"alinx":90,"golng":100}
	)

	txt,err := json.Marshal(number)
	fmt.Printf("%#v %T %#v\n",err,txt,string(txt))

	txt,err = json.Marshal(name)
	fmt.Printf("%#v %T %#v\n",err,txt,string(txt))

	txt,err = json.Marshal(isbody)
	fmt.Printf("%#v %T %#v\n",err,txt,string(txt))

	txt,err = json.Marshal(user)
	fmt.Printf("%#v %T %#v\n",err,txt,string(txt))

	txt,err = json.Marshal(scores)
	fmt.Printf("%#v %T %#v\n",err,txt,string(txt))

	jsonB := `
	{
		"name": "alinx",
		"age": "99",
		"scores": [1,2,3,4,5]
	}
	`
	//任意类型，使用空接口
	var rs map[string]interface{} //值为任意类型是 使用空接口作为值类型
	err = json.Unmarshal([]byte(jsonB), &rs)
	fmt.Printf("%#v %#v",err,rs)

	n,_ := rs["name"]
	fmt.Printf("\n%T\n",n)
}
