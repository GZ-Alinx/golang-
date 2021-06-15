package handers

import (
	"net/http"
	"cmdb/utils"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {
	utils.PaerseTemplate(w, r, []string{"views/user/users.html"}, "users.html", nil)

}
