package LeaderDAO

import "Attendance/Model"

func (leader_DAO *LeaderDAO) View(ApplicationID int, ok bool) (err error) {

	var Application_info Model.Application

	err = leader_DAO.orm.Where("id = ?", ApplicationID).First(&Application_info).Error
	if err != nil {
		return err
	}

	if ok {
		Application_info.Status = 1
	} else {
		Application_info.Status = -1
	}

	err = leader_DAO.orm.Save(&Application_info).Error
	if err != nil {
		return err
	}
	return nil
}
