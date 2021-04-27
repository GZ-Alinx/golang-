package command

import (
	"strconv"
)

var id int = 2

var Commands map[string]func() = map[string]func(){
	// "2": user.Add,
	// "3": user.Modify,
	// "4": user.Delete,
	// "5": user.Query,
} // 指令生成
var Prompts [][2]string = [][2]string{
	// {"2", "添加"},
	// {"3", "修改"},
	// {"4", "删除"},
	// {"5", "查询"},
} // 提示信息

// 注册信息
func Register(desc string, callback func()) {
	Commands[strconv.Itoa(id)] = callback
	Prompts = append(Prompts, [2]string{strconv.Itoa(id), desc})
	id++
}
