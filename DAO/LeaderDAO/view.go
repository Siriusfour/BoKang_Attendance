package LeaderDAO

import (
	"Attendance/Controller/DTO"
	"Attendance/Model"
	"Attendance/Utills"
)

func (leader_DAO *LeaderDAO) View(viewDTO *DTO.ViewDTO) *Utills.MyError {

	//定义一个接受MySQL数据的变量
	var Application_info Model.Application
	var view_user_info Model.User
	var applications_user_info Model.User

	//判断是否是leader，不是的话直接返回错误
	err := leader_DAO.orm.Where("user_id=?", viewDTO.UseID).First(&view_user_info).Error
	if err != nil {
		return Utills.ErrIsDBOperateIsFailed.ErrorAppend(err)
	}

	// 查询 该条记录的申请者的信息，获取到UserID
	err = leader_DAO.orm.Where("id = ?", viewDTO.ApplicationsID).First(&Application_info).Error
	if err != nil {
		return Utills.ErrIsDBOperateIsFailed.ErrorAppend(err)
	}

	// 由UserID查询申请者所在部门
	err = leader_DAO.orm.Where("user_id=?", Application_info.UserID).First(&applications_user_info).Error

	//如果 1.不是leader  2.该leader操作的申请记录不在自己管理的部门  则返回错误
	if view_user_info.Leader == 0 || view_user_info.Leader != applications_user_info.Departmental {
		return Utills.ErrIsUserPermissionDenied
	}

	//根据http标识 通过或驳回
	if viewDTO.Pass {
		Application_info.Status = 1
	} else {
		Application_info.Status = -1
	}

	err = leader_DAO.orm.Save(&Application_info).Error
	if err != nil {
		return Utills.ErrIsDBOperateIsFailed.ErrorAppend(err)
	}
	return nil
}
