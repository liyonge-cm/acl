package resource

import "acl/router"

func init() {
	group := router.NewGroup("resource")
	group.NewRouter("/create", CreateResource)
	group.NewRouter("/getList", GetResourceList)
	group.NewRouter("/get", GetResource)
	group.NewRouter("/update", UpdateResource)
	group.NewRouter("/delete", DeleteResource)
	group.Register()
}
