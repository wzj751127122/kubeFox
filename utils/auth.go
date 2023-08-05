package utils

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func GetClaims(c *gin.Context) (*CustomClaims, error) {
	token := c.Request.Header.Get("token")
	if token == "" {
		return nil, errors.New("请求未携带token,无权限访问")
	}
	// 解析token内容
	claims, err := JwtToken.ParseToken(token)
	if err != nil {
		return nil, err
	}
	return claims, err
}

// func getUser(c *gin.Context) (*model.SysUser, error) {
// 	// 从请求头中获取JWT token
// 	token := c.Request.Header.Get("token")
// 	if token == "" {
// 		return nil, errors.New("请求未携带token,无权限访问")
// 	}
// 	// 解析JWT token
// 	claims, err := JwtToken.ParseToken(token)
// 	if err != nil {
// 		// handle error
// 	}

// 	// 从JWT token中获取用户信息
// 	user := &model.SysUser{
// 		UUID:        claims.UUID,
// 		ID:          claims.ID,
// 		UserName:    claims.Username,
// 		NickName:    claims.NickName,
// 		AuthorityId: claims.AuthorityId,
// 	}

// 	return user,err
// }

// GetUserAuthorityId 从Gin的Context中获取从jwt解析出来的用户角色id
func GetUserAuthorityId(c *gin.Context) (uint, error) {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != nil {
			return 0, err
		} else {
			return cl.AuthorityId, nil
		}
	} else {
		waitUse := claims.(*CustomClaims)
		return waitUse.AuthorityId, nil
	}
}

// GetUserInfo 从Gin的Context中获取从jwt解析出来的用户角色id
func GetUserInfo(c *gin.Context) *CustomClaims {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != nil {
			return nil
		} else {
			return cl
		}
	} else {
		waitUse := claims.(*CustomClaims)
		return waitUse
	}
}
