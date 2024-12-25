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
			Message: fmt.Errorf("Binding ata is Failed:%v", err).Error(),
		})
		return
	}

	err = leader_controller.LeaderServer.View(Request_Message.DTO.(DTO.ViewDTO))
	if err != nil {
		Base.Fail(ctx, Base.Response{
			Message: fmt.Errorf("View Error:%v", err).Error(),
		})
	}

	Base.OK(ctx, Base.Response{
		Data: ViewDTO,
	})

}
