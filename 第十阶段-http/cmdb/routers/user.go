package routers

import (
	"cmdb/handers"
	"net/http"
)

func init() {
	http.HandleFunc("/login/", handers.LoginHander)
	http.HandleFunc("/users/", handers.UserHandler)
}
