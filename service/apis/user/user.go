package user

import "acl/router"

func init() {
	group := router.NewGroup("user")
	group.NewRouter("/create", CreateUser)
	group.NewRouter("/getList", GetUserList)
	group.NewRouter("/get", GetUser)
	group.NewRouter("/update", UpdateUser)
	group.NewRouter("/delete", DeleteUser)
	group.NewRouter("/getUserPermission", GetUserPermission)
	group.Register()
}
