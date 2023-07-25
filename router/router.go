package router

import (
	"k8s-platform/app/opention"
	"k8s-platform/controller"
	"k8s-platform/controller/authory"
	"k8s-platform/controller/menu"
	"k8s-platform/controller/user"
	"k8s-platform/middle"


	"k8s.io/apimachinery/pkg/util/sets"
)

var AlwaysAllowPath sets.String

func InitRouter(opt *opention.Options) {
	apiGroup := opt.GinEngine.Group("/api")
	middle.InstallMiddlewares(apiGroup)
	//安装不需要操作记录路由
	{
		// api.NewApiRouter(apiGroup)
		// operation.NewOperationRouter(apiGroup)
		user.UserinitRoutes(apiGroup)
	}
	// 需要操作记录
	apiGroup.Use(middle.OperationRecord())
	{
		// other.NewSwaggarRoute(apiGroup)
		controller.KubeApiRouter(apiGroup)
		menu.MenuinitRoutes(apiGroup)
		authory.AuthorityinitRoutes(apiGroup)
	}
}
