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
func (My_Base *Base) Application(ctx *gin.Context) {

	//1.====
	var ApplicationsDTO DTO.ApplicationsDTO

	Request_Message := BuildRequest{
		Ctx: ctx,
		DTO: &ApplicationsDTO,
	}

	err := My_Base.Build_request(Request_Message).GetErrors()
	if err != nil {
		ServerFail(My_Base.Ctx, Response{
			Message: Utills.Binding_Data_is_Failed + err.Error(),
		})
		return
	}

	//2.====
	My_Base.IsTokenValid(ctx, Request_Message)

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
