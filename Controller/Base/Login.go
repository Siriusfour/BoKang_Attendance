package Base

import (
	"Attendance/Controller/DTO"
	"Attendance/Utills"
	"fmt"
	"github.com/gin-gonic/gin"
)

// Login 1.把http的消息录入到对象My_Base中
// 2.调用My_Base中的sever层对应的方法
// 3.返回信息（页面，jwt）
func (My_Base *Base) Login(ctx *gin.Context) {

	//1.从ctx获取参数,并绑定到一个dto上
	var loginDTO DTO.LoginDTO

	Request_Message := BuildRequest{
		Ctx: ctx,
		DTO: &loginDTO,
	}

	err := My_Base.Build_request(Request_Message).GetErrors()
	if err != nil {
		ServerFail(ctx, Response{
			Message: fmt.Errorf(Utills.Binding_Data_is_Failed, err).Error(),
		})
		return
	}

	//判断token是否有效
	TokenInfo, TokenErr := IsTokenValid(TokenVerifyInfo{
		UserID:       Request_Message.DTO.(*DTO.LoginDTO).UserID,
		Accesstoken:  Request_Message.DTO.(*DTO.LoginDTO).AccessToken,
		Refreshtoken: Request_Message.DTO.(*DTO.LoginDTO).RefreshToken,
		PassWord:     Request_Message.DTO.(*DTO.LoginDTO).Password,
	})
	if TokenErr != nil {
		ServerFail(ctx, Response{
			Code:    TokenErr.ErrorCode(),
			Message: TokenErr.Error(),
		})
		return
	}

	//调用server层方法
	ApplicationArray, err := My_Base.Server.Login(Request_Message.DTO.(*DTO.LoginDTO))
	if err != nil {
		Fail(ctx, Response{
			Code:    Utills.ErrIsDBOperateIsFailed.ErrorCode(),
			Message: Utills.ErrIsDBOperateIsFailed.ErrorAppend(err.Error()).Error(),
		})

		return
	}

	OK(ctx, Response{
		Message: "Success",
		Data:    ApplicationArray,
		Token:   TokenInfo,
	})

}
