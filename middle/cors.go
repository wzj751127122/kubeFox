package middle

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// func Cors() gin.HandlerFunc {

// 	return func(c *gin.Context) {

// 		//获取请求方法
// 		method := c.Request.Method

// 		//添加跨域响应头
// 		c.Header("Content-Type","application/json")
// 		c.Header("Access-Control-Allow-Origin", "*")
// 		c.Header("Access-Control-Max-Age", "86400")
// 		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
// 		c.Header("Access-Control-Allow-Headers", "X-Token,X-Max,Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
// 		c.Header("Access-Control-Allow-Credentials","false")

// 		if method == "OPTIONS" {
// 			c.AbortWithStatus(http.StatusNoContent)

// 		}

// 		c.Next()

// 	}

// }

// Cores 处理跨域请求，支持options访问
func Cores() gin.HandlerFunc {
	c := cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowHeaders:    []string{"Content-Type", "Access-Token", "Authorization"},
		MaxAge:          6 * time.Hour,
	}

	return cors.New(c)
}
