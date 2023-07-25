package user

import "github.com/gin-gonic/gin"



func UserinitRoutes(ginEngine *gin.RouterGroup) {
	userRoute := ginEngine.Group("/user")
	// user := &userController{}
	{
		userRoute.POST("/login", UserController.Login)
		userRoute.GET("/loginout", UserController.LoginOut)
		userRoute.GET("/getinfo", UserController.GetUserInfo)
		userRoute.PUT("/:id/set_auth", UserController.SetUserAuthority)
		userRoute.DELETE("/:id/delete_user", UserController.DeleteUser)
		userRoute.POST("/:id/change_pwd", UserController.ChangePassword)
		userRoute.PUT("/:id/reset_pwd", UserController.ResetPassword)
	}
}
