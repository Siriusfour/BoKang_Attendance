package Utills

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
	"time"
)

// JwtCustomClaims token的结构体
type JwtCustomClaims struct {
	//通用配置
	jwt.RegisteredClaims
	ID       int
	Password string
}

// 获得用户名密码之后生成token
func Generatetoken(id int, Password string) (string, error) {
	//定义一个自己的JwtCustomClaims类型
	MyJwt := JwtCustomClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			//失效时间
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(viper.GetDuration("Jwt.OutTime") * time.Minute)),
			//颁发时间
			IssuedAt: jwt.NewNumericDate(time.Now()),
			Subject:  "Token",
		},
		ID:       id,
		Password: Password,
	}
	//使用jwt.NewWithClaims对上面的结构体做哈希加密，生成token
	Token := jwt.NewWithClaims(jwt.SigningMethodHS256, MyJwt)

	//使用配置文件里的密钥对token进行签名，并返回一个josn
	return Token.SignedString([]byte(viper.GetString("Jwt.Key")))
}

// ParseJwt 解析token
func ParseJwt(Token string) (*JwtCustomClaims, error) {
	MyJwt := JwtCustomClaims{}
	token, err := jwt.ParseWithClaims(Token, &MyJwt, func(token *jwt.Token) (interface{}, error) {
		return []byte(viper.GetString("Jwt.Key")), nil
	})

	// 检查 token 是否有效
	if err != nil || !token.Valid {
		return nil, errors.New("token is invalid")
	}

	return &MyJwt, nil
}

// IsTokenValid  判断token是否有效
func IsTokenValid(tokenString string) (bool, *JwtCustomClaims) {
	info, err := ParseJwt(tokenString)
	if err != nil {
		return false, nil
	}
	return true, info
}
