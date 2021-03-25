package models

// 首字母大写 可以调用 否则不能调用

// 结构体名 大写 属性大写
// 结构体名称大写 属性小写
type Public struct {
	Private string
	Public  string
}

// 结构体名称小写  属性名称大写
// 结构体名称小写 属性名称小写
type privateStruct struct {
	private string
	Public  string
}

// 函数类型结构体
func NewPrivateStruct() *Public {
	return &Public{}
}
