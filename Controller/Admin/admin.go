package Admin

import (
	"Attendance/Controller/Base"
	"Attendance/Controller/DTO"
	"Attendance/Server"
	"Attendance/Utills"
	"github.com/gin-gonic/gin"
)

type Admin struct {
	Base    Base.Base
	service *Server.AdminServir
}

func (admin Admin) Adduser(context *gin.Context) {

	var AddUserDTO DTO.AddWorkDTO
	option := Base.BuildRequest{
		Ctx: context,
		DTO: &AddUserDTO,
	}
	err := admin.Base.Build_request(option).GetErrors()
	if err != nil {
		return
	}

	admin.service.Admin_AddWork(&AddUserDTO)
	if err != nil {
		Base.ServerFail(admin.Base.Ctx, Base.Response{
			Code:    Utills.Add_Work_fild,
			Message: err.Error(),
		})
		return
	}

	Base.OK(context, Base.Response{
		Message: "ok",
		Data:    AddUserDTO,
	})

}
