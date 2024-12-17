package Base

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
)

// Response 返回给客户端的结构体
type Response struct {
	Staute int `json:"-"`
	//HTTP状态码的扩展，自定义的扩展码，
	Code int `json:"code,omitempty"`
	//本次请求结果的详细描述
	Message string `json:"message,omitempty"`
	//返回的数据
	Data interface{} `json:"data,omitempty"`
	// 请求列表时候返回的总页数
	Sum int64 `json:"sum,omitempty"`
}

// BuildStatus 构筑状态码，优先使用response的状态码，其为空返回code
func BuildStatus(response Response, Code int) int {
	if response.Staute == 0 {
		return Code
	}
	return response.Staute
}

// HttpResponse 设置响应的 JSON 数据和 HTTP 状态码，并向客户端返回默认的状态码
func HttpResponse(ctx *gin.Context, response Response, Status int) {

	//如果结构体为空，终止请求，直接返回状态码
	if reflect.DeepEqual(response, response) {
		ctx.AbortWithStatus(Status)
	}

	//response不为空，将response序列化为json，返回给客户端，并终止本次通话
	ctx.AbortWithStatusJSON(Status, response)

}

// 向处理函数提供的接口，返回失败
func Fail(context *gin.Context, response Response) {
	HttpResponse(context, response, BuildStatus(response, http.StatusBadRequest))
}

func OK(context *gin.Context, response Response) {
	HttpResponse(context, response, BuildStatus(response, http.StatusOK))
}

func ServerFail(context *gin.Context, response Response) {
	HttpResponse(context, response, BuildStatus(response, http.StatusInternalServerError))
}
