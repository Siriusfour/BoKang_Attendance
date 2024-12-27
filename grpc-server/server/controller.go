package server

import (
	"context"
	"encoding/base64"
	"fmt"
	"grpc-server/Utills"
	"grpc-server/proto"
)

// 实现 LoginServiceServer 接口
type Server struct {
	MyProto.UnimplementedLoginServiceServer
}

// 实现 login 方法
func (s *Server) Login(ctx context.Context, req *MyProto.LoginRequest) (*MyProto.LoginResponse, error) {
	// 在这里处理登录逻辑（假设进行简单的密文解密和验证）
	// 这里只是一个简单的示例，实际项目中需要接入数据库等服务

	passwordd, err := base64.StdEncoding.DecodeString(req.Ciphertext)

	password, err := Utills.Decrypt([]byte(Utills.Key), passwordd)
	if err != nil {
		return &MyProto.LoginResponse{}, err
	}

	refreshToken, _ := Utills.Generatetoken(32, string(password))

	if string(password) == "IFNQWFNQWF4F45420DIWK4845KOK779JLL" {
		return &MyProto.LoginResponse{
			UserId:       "123456",
			Username:     "testuser",
			AccessToken:  "access_token_example",
			RefreshToken: refreshToken,
		}, nil
	}
	return nil, fmt.Errorf("invalid credentials")
}

// 实现 refreshToken 方法
func (s *Server) RefreshToken(ctx context.Context, req *MyProto.RefreshTokenRequest) (*MyProto.RefreshTokenResponse, error) {
	// 这里是简单的刷新令牌逻辑，实际项目中需要验证原始刷新令牌等
	if req.RefreshToken == "12345" {
		return &MyProto.RefreshTokenResponse{
			AccessToken: "new_access_token_example",
		}, nil
	}
	return nil, fmt.Errorf("invalid refresh token")
}
