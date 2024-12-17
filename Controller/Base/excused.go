package Base

import (
	"Attendance/Controller/DTO"
	"Attendance/Utills"
	"errors"
	"github.com/gin-gonic/gin"
)

// 1.定义一个DTO，把结构体信息绑定上去
// 2.检查DTO中的token是否有效,请假人信息和jwt内容是否相等
// 3.调用server，在数据库插入待处理请假信息。
// 4.返回ok
func (My_Base *Base) excused(ctx *gin.Context) error {

	//1.====
	var excusedApplicationsDTO DTO.ExcusedApplicationsDTO

	Request_Message := BuildRequest{
		Ctx: ctx,
		DTO: &excusedApplicationsDTO,
	}

	err := My_Base.Build_request(Request_Message).GetErrors()
	if err != nil {
		ServerFail(My_Base.Ctx, Response{
			Message: Utills.Binding_Data_is_Failed + err.Error(),
		})
		return errors.New("binding_data_is_failed:" + err.Error())
	}

	//2.====
	invalid, excused_info := Utills.IsTokenValid(Request_Message.DTO.(DTO.ExcusedApplicationsDTO).Token)
	if invalid {
		Fail(My_Base.Ctx, Response{
			Message: "Token invalid",
		})
	}

	if excused_info.ID != Request_Message.DTO.(DTO.ExcusedApplicationsDTO).UserId {
		Fail(My_Base.Ctx, Response{
			Message: "UserId invalid",
		})
	}

	//3.====
	err = My_Base.Server.Excused(Request_Message.DTO.(DTO.ExcusedApplicationsDTO))
	if err != nil {
		ServerFail(My_Base.Ctx, Response{
			Message: Utills.Insect_is_Failed + "" + err.Error(),
		})
	}

	//4.===
	OK(My_Base.Ctx, Response{
		Message: "success",
	})
	return nil
}
