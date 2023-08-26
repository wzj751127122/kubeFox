package router

import (
	"k8s-platform/app/opention"
	"k8s-platform/controller"
	"k8s-platform/controller/authory"
	"k8s-platform/controller/menu"
	"k8s-platform/controller/operation"
	"k8s-platform/controller/swagger"
	"k8s-platform/controller/user"
	"k8s-platform/middle"

	"github.com/gin-gonic/gin"
	"k8s.io/apimachinery/pkg/util/sets"
)

var AlwaysAllowPath sets.String

func InitRouter(opt *opention.Options) {

	opt.GinEngine.Static("/static", "./static")
	// opt.GinEngine.StaticFS("/", http.Dir("./static"))

	opt.GinEngine.GET("/", func(ctx *gin.Context) {
		// ctx.HTML(http.StatusOK, "index.html", nil)
		ctx.File("static/index.html")
	})
	// 捕获所有路径，返回前端应用的 index.html
	opt.GinEngine.NoRoute(func(c *gin.Context) {
		c.File("static/index.html")
	})
	apiGroup := opt.GinEngine.Group("/api")
	middle.InstallMiddlewares(apiGroup)
	// opt.GinEngine.LoadHTMLFiles("./static/index.html")

	//安装不需要操作记录路由
	{
		// api.NewApiRouter(apiGroup)
		operation.OperationRouter(apiGroup)
		user.UserinitRoutes(apiGroup)
	}
	// 需要操作记录
	apiGroup.Use(middle.OperationRecord())
	{
		swagger.NewSwaggarRoute(apiGroup)
		controller.KubeApiRouter(apiGroup)
		menu.MenuinitRoutes(apiGroup)
		authory.AuthorityinitRoutes(apiGroup)
	}
}
