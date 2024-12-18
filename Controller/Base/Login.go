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

	My_Base.Logger.Info("Login - once")
	//1.从ctx获取参数,并绑定到一个dto上
	var loginDTO DTO.LoginDTO

	Request_Message := BuildRequest{
		Ctx: ctx,
		DTO: &loginDTO,
	}

	err := My_Base.Build_request(Request_Message).GetErrors()
	if err != nil {
		ServerFail(ctx, Response{
			Message: fmt.Errorf("Binding ata is Failed:%v", err).Error(),
		})
		return
	}
	My_User, err := My_Base.Server.Login(Request_Message.DTO.(*DTO.LoginDTO))
	if err != nil {
		Fail(My_Base.Ctx, Response{
			Message: err.Error(),
		})
		return
	}

	//构筑jwt
	token, err := Utills.Generatetoken(My_User.UserID, My_User.Password)

	OK(My_Base.Ctx, Response{

		Message: "登录成功",
		Data: gin.H{
			"token": token,
			"User":  My_User,
		},
	})

}
