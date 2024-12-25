package BaseServer

import (
	"Attendance/Controller/DTO"
	"Attendance/Global"
	"Attendance/Model"
)

func (My_Server *BaseServer) Login(My_DTO DTO.LoginDTO) (*DTO.ApplicationsDTOArray, error) {

	var UserInfo Model.User
	var applicationsDTOArray DTO.ApplicationsDTOArray

	err := Global.DB.Where("user_id=?", My_DTO.UserID).First(&UserInfo).Error
	if err != nil {
		return &applicationsDTOArray, err
	}

	if UserInfo.Leader != 0 {
		err = Global.DB.Where("Department=?", UserInfo.Leader).Find(&applicationsDTOArray.DepartmentApplications).Error
		if err != nil {
			return &applicationsDTOArray, err
		}
	}

	err = Global.DB.Where("user_id=?", UserInfo.UserID).First(&applicationsDTOArray.MyApplications).Error
	if err != nil {
		return &applicationsDTOArray, err
	}

	return &applicationsDTOArray, nil

}
