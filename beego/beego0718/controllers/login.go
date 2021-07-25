package controllers

import "github.com/beego/beego/v2/server/web"

type HomeController struct {
	web.Controller
}

func (c *HomeController) Get() {
	c.TplName = "login.html"
}
