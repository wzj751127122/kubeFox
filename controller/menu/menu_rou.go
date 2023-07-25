package menu

import "github.com/gin-gonic/gin"


func MenuinitRoutes(ginEngine *gin.RouterGroup) {
	menuRoute := ginEngine.Group("/menu")
	// menu:= &menuController{}
	{
		menuRoute.GET("/:authID/getMenuByAuthID", MenuController.GetMenusByAuthID)
		menuRoute.GET("/getBaseMenuTree", MenuController.GetBaseMenus)
		menuRoute.POST("/add_base_menu", MenuController.AddBaseMenu)
		menuRoute.POST("/add_menu_authority", MenuController.AddMenuAuthority)
	}
}
