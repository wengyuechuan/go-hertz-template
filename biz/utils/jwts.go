package utils

import (
	"github.com/dgrijalva/jwt-go/v4"
	"time"
	"uav/biz/conf"
	"uav/biz/consts"
)

// JwtPayLoad jwt中payload数据
type JwtPayLoad struct {
	//Username string `json:"username"`  // 用户名
	NickName string `json:"nick_name"` // 昵称
	Role     int    `json:"role"`      // 权限  1 管理员  2 普通用户  3 游客
	UserID   uint   `json:"user_id"`   // 用户id
	Avatar   string `json:"avatar"`
}

var MySecret []byte

type CustomClaims struct {
	JwtPayLoad
	jwt.StandardClaims
}

// GenToken 创建 Token
func GenToken(user JwtPayLoad) (string, error) {
	MySecret = []byte(conf.Conf.Jwt.Secret)
	claim := CustomClaims{
		user,
		jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(time.Hour * time.Duration(conf.Conf.Jwt.Expires))), // 默认2小时过期
			Issuer:    conf.Conf.Jwt.Issuer,                                                     // 签发人
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString(MySecret)
}

// ParseToken 解析 token
func ParseToken(tokenStr string) (*CustomClaims, error) {
	MySecret = []byte(conf.Conf.Jwt.Secret)
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return MySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, consts.TokenIsInvalid
}
