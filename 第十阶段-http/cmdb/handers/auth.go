package handers

import (
	"net/http"
	"cmdb/utils"
)

func LoginHander(w http.ResponseWriter, r *http.Request) {
	// 打开页面 登录认证
	if r.Method == http.MethodPost {
		// 登录验证  获取页面中获取的数据
		name := r.PostFormValue("name")
		password := r.PostFormValue("password")

		//登录成功后跳转到其他的位置
		//重定向
		if user := services.ValidateLogin(name, password); user != nil {
			http.Redirect(w, r, "/users/", http.StatusFound)
			return
		}
		err = "用户名或密码错误"
	}

	utils.PaerseTemplate(w, r, []string{"view/auth/login.html"}, "login.html", struct {
		Name string
		Err string
	}{Name,Err})
}
