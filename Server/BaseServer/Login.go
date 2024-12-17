package BaseServer

import (
	"Attendance/Controller/DTO"
	"Attendance/Model"
	"errors"
)

func (My_Server *BaseServer) Login(My_DTO DTO.LoginDTO) (Model.User, error) {

	var ErrResult error

	My_user := My_Server.Base_DAO.Get_User_Name_PassWord(My_DTO.Name, My_DTO.PassWord)
	if My_user.ID == 0 {
		ErrResult = errors.New("登录失败，密码或者用户名错误")
	}

	return My_user, ErrResult
}
