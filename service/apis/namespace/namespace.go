package namespace

import "acl/router"

func init() {
	group := router.NewGroup("namespace")
	group.NewRouter("/create", CreateNamespace)
	group.NewRouter("/getList", GetNamespaceList)
	group.NewRouter("/get", GetNamespace)
	group.NewRouter("/update", UpdateNamespace)
	group.NewRouter("/delete", DeleteNamespace)
	group.Register()
}
