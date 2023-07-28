package middle

import (
	"k8s-platform/model"

	"github.com/gin-gonic/gin"

	"k8s.io/apimachinery/pkg/util/sets"
)

var AlwaysAllowPath sets.String

func InstallMiddlewares(ginEngine *gin.RouterGroup) {
	// 初始化可忽略的请求路径
	AlwaysAllowPath = sets.NewString(model.LoginURL, model.LogoutURL, model.WebShellURL)
	ginEngine.Use(Cores(), JWTAuthMiddleware(), GinLogger(), GinRecovery(true))
}
