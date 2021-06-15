package main

import (
	"html/template"
	"os"
)

func main() {
	//tpl,_ := template.New("tpl").ParseFiles("html/header.html", "html/a.html")
	tpl,_ := template.ParseGlob("html/*.html")
	tpl.ExecuteTemplate(os.Stdout, "page.html", struct {
		Title string
	}{"模板文件"})
}
