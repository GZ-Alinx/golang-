package main

import (
	"html/template"
	"os"
)

func main() {
	// 模板定义
	header := `{{ define "header" }}<head><meta charset="utf-8"><title>{{ .Title }}</title></head>{{ end }}`
	page1 := `<!doctype html>
	<html>
		{{ template "header" }}
		<body>
		page1
		</body>
	</html>
	`
	page2 := `<!doctype html>
	<html>
		{{ template "header" }}
		<body>
		page2
		</body>
	</html>`



	tpl,_ := template.New("tpl").Parse(header)
	tpl, _ = tpl.Parse(page1)
	tpl.Execute(os.Stdout, struct {
		Title string
	}{"page1 template定义"})

	tpl2,_ := template.New("tpl2").Parse(header)
	tpl2, _ = tpl2.Parse(page2)
	tpl2.Execute(os.Stdout, struct {
		Title string
	}{"page2 template定义"})

}
