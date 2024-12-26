package Base

import (
	"Attendance/Global"
	"Attendance/Server/BaseServer"
	"Attendance/Utills"
	"Attendance/grpc/MyProto"
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"reflect"
	"strconv"
	"time"
)

// Base         Leader 和  worker 通用结构体与方法
type Base struct {
	Ctx    *gin.Context
	errors error
	Logger *zap.SugaredLogger
	Server *BaseServer.BaseServer
}

// BuildRequest http 请求体结构
type BuildRequest struct {
	Ctx *gin.Context
	DTO any
}

type TokenVerifyInfo struct {
	UserID        int
	Access_token  string
	Refresh_token string
	PassWord      string
}

func NewBase() *Base {
	return &Base{
		Logger: Global.Logger,
		Server: BaseServer.New_Base_Server(),
	}

}

func (My_Base *Base) AddError(NewErr error) {

	My_Base.errors = Utills.AppendErr(My_Base.errors, NewErr)
}

func (My_Base *Base) GetErrors() error {
	return My_Base.errors
}

// Build_request 从HTTP中提取信息， 构筑请求结构体
func (My_Base *Base) Build_request(Request BuildRequest) *Base {

	var ErrRuslt error

	My_Base.Ctx = Request.Ctx

	if Request.DTO != nil {

		ErrRuslt = My_Base.Ctx.ShouldBindBodyWithJSON(Request.DTO)

		if ErrRuslt != nil {
			Err := parseErr(ErrRuslt, Request.DTO)
			My_Base.AddError(Err)
			Fail(My_Base.Ctx, Response{
				Message: My_Base.GetErrors().Error(),
			})
		}
	}
	return My_Base
}

// 解析验证失败的err，处理成可供客户端使用的错误提示信息
func parseErr(Errs error, tager interface{}) error {

	var ErrRuslt error
	var ErrValidation validator.ValidationErrors
	ok := errors.As(Errs, &ErrValidation)
	if !ok {
		return Errs
	}

	//返回指针类型所指向对象的类型
	Flieds := reflect.TypeOf(tager).Elem()
	for _, FliedsErr := range ErrValidation {
		Flied, _ := Flieds.FieldByName(FliedsErr.Field())
		ErrMessageTag := fmt.Sprintf("%s_err", FliedsErr.Tag())
		ErrMessage := Flied.Tag.Get(ErrMessageTag)
		if ErrMessage == "" {
			ErrMessage = Flied.Tag.Get("Message")
		}

		if ErrMessage == "" {
			ErrMessage = fmt.Sprintf("%s:%s Err", FliedsErr.Field(), FliedsErr.Tag())
		}

		ErrRuslt = Utills.AppendErr(ErrRuslt, errors.New(ErrMessage))
	}

	return ErrRuslt
}

// IsTokenValid 验证token
func (My_Base *Base) IsTokenValid(tokenVerifyInfo TokenVerifyInfo) (TokenVerifyInfo, *Utills.MyError) {

	redis_AccseeToken := strconv.Itoa(tokenVerifyInfo.UserID) + "_AccseeToken"
	redis_RefreshToken := strconv.Itoa(tokenVerifyInfo.UserID) + "_RefreshToken"

	LoginTokenInfo := TokenVerifyInfo{}
	switch {

	//http请求里只有Access_Token
	case tokenVerifyInfo.Access_token != "":
		{
			ak, err := Global.RedisClient.Get(context.Background(), redis_AccseeToken).Result()
			if err != nil || ak == "" || ak != tokenVerifyInfo.Access_token {

				return LoginTokenInfo, Utills.ErrIsATokenIsInvalid
			}
		}

	//http请求里只有Refresh_Token
	case tokenVerifyInfo.Refresh_token != "":
		{
			rk, err := Global.RedisClient.Get(context.Background(), redis_RefreshToken).Result()
			if (err != nil || rk == "") || rk != tokenVerifyInfo.Refresh_token {
				return LoginTokenInfo, Utills.ErrIsRTokenIsInvalid
			}

			ak, err := Global.Grpc_Client.RefreshToken(context.Background(), &MyProto.RefreshTokenRequest{RefreshToken: rk})

			err = Global.RedisClient.Set(context.Background(), redis_AccseeToken, ak, time.Duration(viper.GetInt("key.Refresh_Token_OutTime"))*10080*60).Err()

			//return LoginTokenInfo{Access_token: ak.AccessToken}, nil
			LoginTokenInfo.Access_token = ak.AccessToken

		}

		//http请求里只有Password
	case tokenVerifyInfo.PassWord != "":
		{

			//对称加密
			key := []byte(viper.GetString("key.privateKey"))
			cipherText, err := Utills.Encrypt(key, []byte(tokenVerifyInfo.PassWord))
			if err != nil {

				return LoginTokenInfo, Utills.NewMyError(err.Error(), Utills.PassWordCryptoIsFailed)
			}

			encodedCipherText := base64.StdEncoding.EncodeToString(cipherText)
			println(encodedCipherText)

			var LoginRequest MyProto.LoginRequest

			LoginRequest.Ciphertext = encodedCipherText

			//调用rpc验证密码
			LoginResponse, err := Global.Grpc_Client.Login(context.Background(), &LoginRequest)
			if err != nil {
				return LoginTokenInfo, Utills.NewMyError(err.Error(), Utills.PassWordVerifyIsFailed)
			}

			//如果密码正确,把刷新的ak、rk写入redis
			if LoginResponse != nil {
				Global.RedisClient.Set(context.Background(), LoginResponse.UserId+"_AccseeToken", LoginResponse.AccessToken, time.Duration(viper.GetInt("key.Access_Token_OutTime")))
				Global.RedisClient.Set(context.Background(), LoginResponse.UserId+"_RefreshToken", LoginResponse.AccessToken, time.Duration(viper.GetInt("key.Refresh_Token_OutTime")))

				LoginTokenInfo.Access_token = LoginResponse.AccessToken
				LoginTokenInfo.Refresh_token = LoginResponse.RefreshToken

			} else {
				return LoginTokenInfo, Utills.NewMyError(err.Error(), Utills.PassWordVerifyIsFailed)
			}
		}

	}

	return LoginTokenInfo, nil
}
