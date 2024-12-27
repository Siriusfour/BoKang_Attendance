package Leader

import (
	"Attendance/Controller/Base"
	"Attendance/Controller/DTO"
	"fmt"
	"github.com/gin-gonic/gin"
)

func (leader_controller *LeaderController) View(ctx *gin.Context) {

	//1.从ctx获取参数,并绑定到一个dto上
	var ViewDTO DTO.ViewDTO

	Request_Message := Base.BuildRequest{
		Ctx: ctx,
		DTO: &ViewDTO,
	}

	//从http绑定数据
	err := leader_controller.BaseController.Build_request(Request_Message).GetErrors()
	if err != nil {
		Base.ServerFail(ctx, Base.Response{
			Message: fmt.Errorf("Binding data is Failed:%v", err).Error(),
		})
		return
	}

	//token鉴权
	_, TokenErr := Base.IsTokenValid(Base.TokenVerifyInfo{
		Accesstoken: Request_Message.DTO.(*DTO.ViewDTO).AccessToken,
		UserID:      Request_Message.DTO.(*DTO.ViewDTO).UseID,
	})
	if TokenErr != nil {
		Base.Fail(ctx, Base.Response{
			Code:    TokenErr.Code,
			Message: TokenErr.Error(),
		})
		return
	}

	ViewErr := leader_controller.LeaderServer.View(Request_Message.DTO.(*DTO.ViewDTO))
	if ViewErr != nil {
		Base.Fail(ctx, Base.Response{
			Message: ViewErr.Error(),
			Code:    ViewErr.Code,
		})
		return
	}

	Base.OK(ctx, Base.Response{
		Message: "Success",
	})

}
