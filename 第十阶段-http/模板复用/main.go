package main

import (
	"html/template"
	"log"
	"os"
)

type User struct {
	ID string
	Name string
}

func main() {
	//模板复用，重写模板中的某块内容 为内容当以一个逻辑块 {{block "name"}} xxx {{ end }}

	// desc 是模板的名字
	txt := `我是 {{ block "desc" . }} {{ . }} {{ end }}`
	data := User{"1","李雄"}
	tpl, err := template.New("overwrite").Parse(txt) // 解析模板
	if err != nil {
		log.Fatal(err)
	}

	//重写模板 覆盖它 desc 是模板名字
	overwirteTxt := `{{ define "desc" }} {{ . }}, 欢迎你 {{ end }}`
	tpl,err = tpl.Parse(overwirteTxt)
	if err != nil {
		log.Fatal(err)
	}
	tpl.Execute(os.Stdout, data)
}

