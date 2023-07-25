package authory

import (


	"github.com/gin-gonic/gin"
)

func AuthorityinitRoutes(ginEngine *gin.RouterGroup) {
	casRoute := ginEngine.Group("/authority")
	{
		casRoute.GET("/getPolicyPathByAuthorityId", AuthoritiyController.GetPolicyPathByAuthorityId)
		casRoute.POST("/updateCasbinByAuthority", AuthoritiyController.UpdateCasbinByAuthorityId)
		casRoute.GET("/getAuthorityList", AuthoritiyController.GetAuthorityList)
	}
}

