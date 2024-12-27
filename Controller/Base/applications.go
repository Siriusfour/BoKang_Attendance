package Base

import (
	"Attendance/Controller/DTO"
	"Attendance/Utills"
	"github.com/gin-gonic/gin"
)

// Application 1.定义一个DTO，把结构体信息绑定上去
// 2.检查DTO中的token是否有效,请假人信息和jwt内容是否相等 (grpc远程判断)
// 3.调用server，在数据库插入待处理请假信息。
// 4.返回ok
func (My_Base *BaseController) Application(ctx *gin.Context) {

	//1.====
	var ApplicationsDTO DTO.ApplicationsDTO

	Request_Message := BuildRequest{
		Ctx: ctx,
		DTO: &ApplicationsDTO,
	}

	err := My_Base.Build_request(Request_Message)
	if err != nil {

		Fail(ctx, Response{
			Message: Utills.ErrIsBindingDataIsFailed.ErrorAppend(err).Error(),
			Code:    Utills.ErrIsBindingDataIsFailed.Code,
		})

		return
	}

	_, TokenErr := IsTokenValid(TokenVerifyInfo{
		UserID:      Request_Message.DTO.(*DTO.ApplicationsDTO).UserID,
		Accesstoken: Request_Message.DTO.(*DTO.ApplicationsDTO).AccessToken,
	})
	if TokenErr != nil {
		Fail(My_Base.Ctx, Response{
			Message: TokenErr.Error(),
			Code:    TokenErr.ErrorCode(),
		})
		return
	}

	//3.====
	err = My_Base.Server.Application(Request_Message.DTO.(*DTO.ApplicationsDTO))
	if err != nil {
		ServerFail(My_Base.Ctx, Response{
			Message: Utills.Insect_is_Failed + "" + err.Error(),
		})
		return
	}

	//4.===
	OK(My_Base.Ctx, Response{
		Message: "success",
	})
	return
}
