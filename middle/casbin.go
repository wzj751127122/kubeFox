package middle

import (
	"fmt"

	"k8s-platform/logic"
	"k8s-platform/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CasbinHandler 拦截器
func CasbinHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		if AlwaysAllowPath.Has(c.Request.URL.Path) {
			return
		}
		waitUse, err := utils.GetClaims(c)
		if err != nil {
			ResponseError(c,CodeServerBusy)
			c.Abort()
			return
		}
		// 获取请求的PATH
		obj := c.Request.URL.Path
		// 获取请求方法
		act := c.Request.Method
		// 获取用户的角色
		sub := strconv.Itoa(int(waitUse.AuthorityId))
		e := logic.Casbin() // 判断策略中是否存在
		success, _ := e.Enforce(sub, obj, act)
		if success {
			c.Next()
		} else {
			ResponseErrorWithMsg(c,CodeServerBusy, fmt.Errorf("角色ID %d 请求 %s %s 无权限", waitUse.AuthorityId, act, obj))
			c.Abort()
			return
		}
	}
}