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

	AccessTokenIsInvalid = 1001
	RefreshTokenIsValid  = 1002
	PassWordIsInvalid    = 1003

	PassWordCryptoIsFailed = 1005
	PassWordVerifyIsFailed = 1006

	//数据库操作，2xxx

	DB_operateIsFailed      = 2001
	DB_UserPermissionDenied = 2002
)

var (
	//登录失败
	ErrIsATokenIsInvalid   = NewMyError("access_token is invalid", AccessTokenIsInvalid)
	ErrIsRTokenIsInvalid   = NewMyError("refresh_token is invalid", RefreshTokenIsValid)
	ErrIsPassWordIsInvalid = NewMyError("password is invalid", PassWordIsInvalid)

	//数据库操作错误
	ErrIsDBOperateIsFailed    = NewMyError("DB_Operate_Is_Failed", DB_operateIsFailed)
	ErrIsUserPermissionDenied = NewMyError("用户权限不足：", DB_UserPermissionDenied)
)

// 定义一个自定义错误类型,有错误信息和错误码
type MyError struct {
	Message string // 错误消息
	Code    int    // 错误代码
}

// 实现 error 接口的 Error 方法
func (e *MyError) Error() string {
	return e.Message
}

func (e *MyError) ErrorCode() int {
	return e.Code
}

// 将err的错误描述添加到该自定义错误
func (e *MyError) ErrorAppend(errMsg string) *MyError {
	NewErr := *e
	NewErr.Message = NewErr.Message + ": " + errMsg
	return &NewErr
}

// 创建一个构造函数，用于初始化自定义错误
func NewMyError(message string, code int) *MyError {
	return &MyError{
		Message: message,
		Code:    code,
	}
}
