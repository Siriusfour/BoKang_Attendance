package Utills

import (
	"errors"
	"time"
)
import "github.com/golang-jwt/jwt/v4"

// // JwtCustomClaims token的结构体
type JwtCustomClaims struct {
	//通用配置
	jwt.RegisteredClaims
	ID       int
	Password string
}

func Generatetoken(id int, Password string) (string, error) {
	//定义一个自己的JwtCustomClaims类型
	MyJwt := JwtCustomClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			//失效时间
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(AccseeToken_OutTime * time.Minute)),
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
	return Token.SignedString([]byte(Key))
}

// ParseJwt 解析token
func ParseJwt(Token string) (*JwtCustomClaims, error) {
	MyJwt := JwtCustomClaims{}
	token, err := jwt.ParseWithClaims(Token, &MyJwt, func(token *jwt.Token) (interface{}, error) {
		return []byte(Key), nil
	})

	// 检查 token 是否有效
	if err != nil || !token.Valid {
		return nil, errors.New("token is invalid")
	}

	return &MyJwt, nil
}
