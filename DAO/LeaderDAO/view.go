package LeaderDAO

import "Attendance/Model"

func (leader_DAO *LeaderDAO) View_Access(ApplicationID int) {

	var Application_info Model.Application

	leader_DAO.orm.Where("id = ?", ApplicationID).First(&Application_info)

	Application_info.Status = 1

	leader_DAO.orm.Save(&Application_info)
}

func (leader_DAO *LeaderDAO) View_rejected(ApplicationID int) {

	var Application_info Model.Application

	leader_DAO.orm.Where("id = ?", ApplicationID).First(&Application_info)

	Application_info.Status = -1

	leader_DAO.orm.Save(&Application_info)
}
