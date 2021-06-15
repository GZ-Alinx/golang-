package main


import (
	"log"
	"os"
	"strings"
	"text/template"
)



type User struct {
	 Id int
	 Name string
}

func main() {
	// 定义模板字符串
	// 占位符
	// 显示数据  {{ . }} .传递的数据
	//txt := `{{ if . }}
	//我是李雄
	//{{ else }}
	//alinx
	//{{end}}`


	////模板函数拿到特定的索引的值
	//txt := `{{ index . 2 }}`
	//data := []int{1,2,3,4,5}


	//// 遍历数据
	//// for 中 . 表示遍历的每个元素 key value 获取 或遍历单个值
	//txt := `{{ range $i, $v := .  }}
	//	{{ $i }}:{{ $v }}
	//{{ end }}`
	//data := []int{1,2,3,4,5}

	// 遍历
	//txt := `{{ range . }}
	//{{ . }}
	//{{ end }}
	//`
	//txt := `{{ range $v := . }}
	//{{ $v }}
	//{{ end }}
	//`
	//txt := `{{ range $k, $v := . }}
	//{{ $k }}:{{ $v }}
	//{{ end }}
	//`
	//data := map[string]string{"lx":"贵州","调度":"system"}





	// 只想获取数据某个属性 特定的值
	//txt := `登录: {{ .Name }}`
	//data := User{1, "alinx"}




	// 生成表格 遍历生成
//	data := []User{{1,"贵州lx1"},{2,"贵州lx2"},{3,"贵州lx3"},{4,"贵州lx4"}}
//	txt := `
//	<table>
//		<thead>
//			<tr>
//				<th>Id</th>
//				<th>Name</th>
//			</tr>
//		</thead>
//		<tbody>
//		{{ range . }}
//			<tr>
//				<td>{{ .Id}}</td>
//				<td>{{- .Name -}}</td>
//			</tr>
//		{{ end }}
//		</tbody>
//	</table>
//`
	//去掉空行
	//{{- .Name -}}


	//data := "alinx"
	//data := []int{1,2,3,4,5,6}
	//data := []int{1,2,3}

	//data := true
	// data true 显示李雄 false 显示 alinx

	//data := map[string]string{"alinx": "贵州"}
	//data := []User{{1, "alinx"},{2,"lx"}}




	// 模板内置的函数 {text/html}/template 用法一致
	/*
	文档地址: https://golang.google.cn/pkg/html/template/
		index 索引
		len 获取长度
		or not 和go语言一样
		print 打印
		urlqyery url转码


	*/


	// 自定义函数 upper lower 大写小写
	txt := `{{ . }} {{ upper . }} {{ lower . }}`
	data := "Abc"
	funcs := template.FuncMap{
		"upper": func(txt string) string {
			return strings.ToUpper(txt)
		},
		"lower": strings.ToLower,
		"tl": func(txt string, isLower bool) string {
			if isLower {
				return strings.ToLower(txt)
			}
			return strings.ToUpper(txt)
		},
	}


	// 解析模板
	tpl := template.New("txtTemplate")
	tpl = tpl.Funcs(funcs) // 指定自定义函数
	tpl, err := tpl.Parse(txt)
	if err != nil {
		log.Fatal(err)
	}

	// 执行输出到控制台
	err = tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatal(err)
	}
}
