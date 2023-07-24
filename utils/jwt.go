package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
	uuid "github.com/satori/go.uuid"

	"github.com/wonderivan/logger"
)

var JwtToken jwttoken

// CustomClaims 自定义声明类型 并内嵌jwt.RegisteredClaims
// jwt包自带的jwt.RegisteredClaims只包含了官方字段
// 假设我们这里需要额外记录一个username字段，所以要自定义结构体
// 如果想要保存更多信息，都可以添加到这个结构体中
//
//	type CustomClaims struct {
//		// 可根据需要自行添加字段
//		Username             string `json:"username"`
//		Password             int64  `json:"password"`
//		jwt.RegisteredClaims        // 内嵌标准的声明
//	}
type jwttoken struct {
	secret string
}

func RegisterJwt(secret string) {
	JwtToken.secret = secret
}

type BaseClaims struct {
	UUID        uuid.UUID
	ID          int
	Username    string
	NickName    string
	AuthorityId uint
}

// CustomClaims 自定义token中携带的信息
type CustomClaims struct {
	BaseClaims
	jwt.StandardClaims
}

// const TokenExpireDuration = time.Hour * 24

// CustomSecret 用于加盐的字符串
// var CustomSecret = []byte("夏天夏天悄悄过去")

// // GenToken 生成JWT
// func GenToken(userID int64, username string) (string, error) {
// 	// 创建一个我们自己的声明
// 	claims := CustomClaims{
// 		UserID:   userID,
// 		Username: "username", // 自定义字段
// 		RegisteredClaims: jwt.RegisteredClaims{
// 			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(viper.GetInt("auth.jwt_expire")) * time.Hour)),
// 			Issuer:    "bluebell", // 签发人
// 		},
// 	}
// 	// 使用指定的签名方法创建签名对象
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	// 使用指定的secret签名并获得完整的编码后的字符串token
// 	return token.SignedString(CustomSecret)
// }

func (j *jwttoken) GenerateToken(baseClaims BaseClaims) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(24 * time.Hour)
	claims := CustomClaims{
		baseClaims,
		jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000, // 签名生效时间
			ExpiresAt: expireTime.Unix(),
			Issuer:    "kubeFox", // 签名的发行者
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString([]byte(j.secret))
	return token, err
}

// ParseToken 解析JWT
func (j *jwttoken) ParseToken(tokenString string) (claims *CustomClaims, err error) {
	// 解析token
	// 如果是自定义Claim结构体则需要使用 ParseWithClaims 方法

	var mc = new(CustomClaims)
	token, err := jwt.ParseWithClaims(tokenString, mc, func(token *jwt.Token) (i interface{}, err error) {
		// 直接使用标准的Claim则可以直接使用Parse方法
		//token, err := jwt.Parse(tokenString, func(token *jwt.Token) (i interface{}, err error) {
		// return CustomSecret, nil
		return []byte(j.secret), nil
	})
	if err != nil {

		logger.Error("parse token failed", err)
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errors.New("TokenMalformed")
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, errors.New("TokenExpired")
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, errors.New("TokenNotValidYet")
			} else {
				return nil, errors.New("TokenInvalid")
			}

		}
		// return nil, err
	}
	// 对token对象中的Claim进行类型断言
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
