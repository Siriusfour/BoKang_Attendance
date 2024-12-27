package BaseServer

import (
	"Attendance/Controller/DTO"
	"Attendance/Model"
	"gorm.io/gorm"
	"time"
)

func (MyServer *BaseServer) Application(My_DTO *DTO.ApplicationsDTO) error {

	Application := Model.Application{
		gorm.Model{
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: gorm.DeletedAt{},
		},
		My_DTO.Name,
		My_DTO.UserID,
		My_DTO.Message,
		My_DTO.StartTime,
		My_DTO.EndTime,
		My_DTO.Department,
		My_DTO.Leave_type,
		0,
	}

	err := MyServer.BaseDAO.Application(Application, *My_DTO.ApplicationID)
	if err != nil {
		return err
	}

	return nil
}
