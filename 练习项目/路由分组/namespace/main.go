package main

import (
	"namespace/controllers/roles"
	"namespace/controllers/users"
	_ "namespace/tasks"

	"github.com/beego/beego/v2/server/web"
)

func main() {
	userNamespace := web.NewNamespace("/users",
		web.NSAutoRouter(new(users.IndexController)),
	)
	rolesNamespace := web.NewNamespace("/roles",
		web.NSAutoRouter(new(roles.IndexController)),
	)

	web.AddNamespace(userNamespace)
	web.AddNamespace(rolesNamespace)
	web.Run()
}
