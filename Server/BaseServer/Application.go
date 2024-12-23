package BaseServer

import (
	"Attendance/Controller/DTO"
	"Attendance/Model"
)

func (My_Server *BaseServer) Application(My_DTO *DTO.ApplicationsDTO) error {

	Application := Model.Application{
		My_DTO.Name,
		My_DTO.UserId,
		My_DTO.Message,
		My_DTO.StartTime,
		My_DTO.EndTime,
		My_DTO.Department,
		My_DTO.Leave_type,
		1,
	}

	err := My_Server.Base_DAO.Application(Application)
	if err != nil {
		return err
	}

	return nil
}
