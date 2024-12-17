package Base

import (
	"Attendance/Global"
	"Attendance/Server/BaseServer"
	"Attendance/Utills"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"reflect"
)

// Base         Admin 和  worker 通用结构体与方法
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

// 从HTTP中提取信息， 构筑请求结构体
func (My_Base *Base) Build_request(Request BuildRequest) *Base {

	var ErrRuslt error
	My_Base.Ctx = Request.Ctx

	if Request.DTO != nil {

		ErrRuslt = My_Base.Ctx.ShouldBind(Request.DTO)

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
