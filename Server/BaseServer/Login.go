package BaseServer

import (
	"Attendance/Controller/DTO"
	"Attendance/Global"
	"Attendance/Model"
)

func (My_Server *BaseServer) Login(My_DTO *DTO.LoginDTO) (*Model.ApplicationsArray, error) {

	var UserInfo Model.User
	var ApplicationsArray Model.ApplicationsArray

	var MyApplications []Model.Application
	var DepartmentApplications []Model.Application

	//查询
	err := Global.DB.Where("user_id=?", My_DTO.UserID).First(&UserInfo).Error
	if err != nil {
		return &ApplicationsArray, err
	}

	if UserInfo.Leader != 0 {
		err = Global.DB.Where("Department=?", UserInfo.Leader).Find(&MyApplications).Error
		if err != nil {
			return &ApplicationsArray, err
		}
	}

	err = Global.DB.Where("user_id=?", UserInfo.UserID).Find(&DepartmentApplications).Error
	if err != nil {
		return &ApplicationsArray, err
	}

	ApplicationsArray.MyApplications = MyApplications
	ApplicationsArray.DepartmentApplications = DepartmentApplications

	return &ApplicationsArray, nil

}
