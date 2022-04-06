package main


import （
  "github.com/tidwall/gjson"
）


// 下载包 go get github.com/tidwall/gjson


func main() {
  jsondata := `{"name":{"first":"123", "last":"456"}, "age":"44"}`
  lastName := gjson.Get(jsondata, "name.last")
  fmt.Println("last name", lastName.String())
  
  
  age := gjson.Get(jsondata, "age")
  fmt.Println("age:", age.Int())
}
