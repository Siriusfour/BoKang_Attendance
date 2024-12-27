package BaseDAO

import (
	"Attendance/Model"
	"time"
)

func (My_BaseDAO *BaseDAO) Application(My_Mode Model.Application, ApplicationID int) error {

	var mode Model.Application

	//若传入的id是0，则说明是新的申请
	if ApplicationID == 0 {
		result := My_BaseDAO.orm.Create(&My_Mode)
		if result.Error != nil {
			return result.Error
		}
		return nil

		//否则为旧申请，直接在库里找到记录再修改
	} else {
		result := My_BaseDAO.orm.Where("ID=?", ApplicationID).First(&mode)
		if result.Error != nil {
			return result.Error
		}

		My_Mode.ID = mode.ID
		My_Mode.CreatedAt = mode.CreatedAt
		My_Mode.UpdatedAt = time.Now()

		mode = My_Mode
		My_BaseDAO.orm.Save(&mode)
	}

	return nil
}
