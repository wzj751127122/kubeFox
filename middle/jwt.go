package middle

import (
	"k8s-platform/utils"

	"github.com/gin-gonic/gin"
)

// JWTAuthMiddleware 基于JWT的认证中间件
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 这里假设Token放在Header的Authorization中，并使用Bearer开头
		// 这里的具体实现方式要依据你的实际业务情况决定

		//对登录接口放行
		if len(c.Request.URL.String()) >= 10 && c.Request.URL.String()[0:10] == "api/login" {
			c.Next()
		}
		// } else {
		// 	authHeader := c.Request.Header.Get("Authorization")
		// 	if authHeader == "" {
		// 		c.JSON(http.StatusBadRequest, gin.H{
		// 			"msg":  "请求未携带Token，无权限访问",
		// 			"data": nil,
		// 		})
		// 		c.Abort()
		// 		return
		// 	}
		// 按空格分割(如果授权方式是Bearer的话使用这里)
		// parts := strings.SplitN(authHeader, " ", 2)
		// if !(len(parts) == 2 && parts[0] == "Bearer") {
		// 	controllers.ResponseError(c, controllers.CodeInvalidToken)
		// 	c.Abort()
		// 	return
		// }
		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		// mc, err := jwt.ParseToken(parts[1])
		// claims, err := utils.JwtToken.ParseToken(authHeader)
		// if err != nil {
		// 	if err.Error() == "TokenExpired" {
		// 		c.JSON(http.StatusBadRequest, gin.H{
		// 			"msg":  "授权已过期",
		// 			"data": nil,
		// 		})
		// 		c.Abort()
		// 		return
		// 	}
		// 处理验证逻辑
		claims, err := utils.GetClaims(c)
		if err != nil {
			ResponseErrorWithMsg(c, CodeInvalidToken, err)
			c.Abort()
			return
		}
		c.Abort()
		return

		// 将当前请求的username信息保存到请求的上下文c上
		
		c.Set("claims", claims)
		c.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
	}
}
