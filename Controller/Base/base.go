package Base

import (
	"Attendance/Global"
	"Attendance/Server/BaseServer"
	"Attendance/Utills"
	"Attendance/grpc/MyProto"
	"context"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"strconv"
	"time"
)

// BaseController         Leader 和  worker 通用结构体与方法
type BaseController struct {
	Ctx    *gin.Context
	errors Utills.MyError
	Logger *zap.SugaredLogger
	Server *BaseServer.BaseServer
}

// BuildRequest http 请求体结构
type BuildRequest struct {
	Ctx *gin.Context
	DTO any
}

type TokenVerifyInfo struct {
	UserID       int
	Accesstoken  string
	Refreshtoken string
	PassWord     string
}

func NewBase() *BaseController {
	return &BaseController{
		Logger: Global.Logger,
		Server: BaseServer.New_Base_Server(),
	}

}

// Build_request 从HTTP中提取信息， 构筑请求结构体
func (My_Base *BaseController) Build_request(Request BuildRequest) error {

	My_Base.Ctx = Request.Ctx

	if Request.DTO != nil {

		err := My_Base.Ctx.ShouldBindBodyWithJSON(Request.DTO)
		if err != nil {
			return err
		}
	}
	return nil
}

// IsTokenValid 验证token
func IsTokenValid(tokenVerifyInfo TokenVerifyInfo) (TokenVerifyInfo, *Utills.MyError) {

	redis_AccseeToken := strconv.Itoa(tokenVerifyInfo.UserID) + "_AccessToken"
	redis_RefreshToken := strconv.Itoa(tokenVerifyInfo.UserID) + "_RefreshToken"

	LoginTokenInfo := TokenVerifyInfo{}
	switch {

	//http请求里只有Access_Token
	case tokenVerifyInfo.Accesstoken != "":
		{
			ak, err := Global.RedisClient.Get(context.Background(), redis_AccseeToken).Result()
			if err != nil || ak == "" || ak != tokenVerifyInfo.Accesstoken {

				return LoginTokenInfo, Utills.ErrIsATokenIsInvalid
			}
		}

	//http请求里只有Refresh_Token
	case tokenVerifyInfo.Refreshtoken != "":
		{
			rk, err := Global.RedisClient.Get(context.Background(), redis_RefreshToken).Result()
			if (err != nil || rk == "") || rk != tokenVerifyInfo.Refreshtoken {
				return LoginTokenInfo, Utills.ErrIsRTokenIsInvalid
			}

			ak, err := Global.Grpc_Client.RefreshToken(context.Background(), &MyProto.RefreshTokenRequest{RefreshToken: rk})

			err = Global.RedisClient.Set(context.Background(), redis_AccseeToken, ak, time.Duration(viper.GetInt("key.RefreshToken_OutTime"))).Err()

			//return LoginTokenInfo{Access_token: ak.AccessToken}, nil
			LoginTokenInfo.Accesstoken = ak.AccessToken

		}

		//http请求里只有Password
	case tokenVerifyInfo.PassWord != "":
		{

			//对称加密
			key := []byte(viper.GetString("key.privateKey"))
			cipherText, err := Utills.Encrypt(key, []byte(tokenVerifyInfo.PassWord))
			if err != nil {

				return LoginTokenInfo, Utills.ErrIsPassWordCryptoIsFailed.ErrorAppend(err)
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

				LoginTokenInfo.Accesstoken = LoginResponse.AccessToken
				LoginTokenInfo.Refreshtoken = LoginResponse.RefreshToken

			} else {
				return LoginTokenInfo, Utills.ErrIsPassWordIsInvalid
			}
		}

	}

	return LoginTokenInfo, nil
}
