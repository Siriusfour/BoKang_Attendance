package Utills

import (
	"fmt"
)

func AppendErr(existing error, newErr error) error {
	if existing == nil {
		return newErr
	}
	return fmt.Errorf("%v; %w", existing, newErr)
}

const (
	Binding_Data_is_Failed = "binding data is failed:"
	Insect_is_Failed       = "insect is failed"
	Access_Token_is_failed = "access token is failed"
	Query_is_failed        = "query userinfo is failed"

	AccessTokenIsInvalid   = 1001
	RefreshTokenIsValid    = 1002
	PassWordIsInvalid      = 1003
	QueryIsFailed          = 1004
	PassWordCryptoIsFailed = 1005
	PassWordVerifyIsFailed = 1006

	LoginDTO    = 1
	Application = 2
)

var (
	ErrIsATokenIsInvalid   = NewMyError("access_token is invalid", AccessTokenIsInvalid)
	ErrIsRTokenIsInvalid   = NewMyError("refresh_token is invalid", RefreshTokenIsValid)
	ErrIsPassWrodIsInvalid = NewMyError("password is invalid", PassWordIsInvalid)
)

// 定义一个自定义错误类型
type MyError struct {
	Message string // 错误消息
	Code    int    // 错误代码
}

// 实现 error 接口的 Error 方法
func (e *MyError) Error() string {
	return fmt.Sprintf("Code: %d, Message: %s", e.Code, e.Message)
}

func (e *MyError) ErrorCode() int {
	return e.Code
}

// 创建一个构造函数，用于初始化自定义错误
func NewMyError(message string, code int) *MyError {
	return &MyError{
		Message: message,
		Code:    code,
	}
}
