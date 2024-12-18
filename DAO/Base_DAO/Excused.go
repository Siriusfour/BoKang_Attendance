package Base_DAO

import (
	"Attendance/Model"
)

func (My_BaseDAO *BaseDAO) Excused(My_Mode Model.Pending_Application) error {

	result := My_BaseDAO.orm.Create(My_Mode)

	if result.Error != nil {
		return result.Error
	}
	return nil

}
