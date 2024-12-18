package BaseServer

import (
	"Attendance/Controller/DTO"
	"Attendance/Model"
)

func (My_Server *BaseServer) Excused(My_DTO DTO.ExcusedApplicationsDTO) error {

	Pending_Application := Model.Pending_Application{
		My_DTO.Name,
		My_DTO.UserId,
		My_DTO.Message,
		My_DTO.StartTime,
		My_DTO.EndTime,
		My_DTO.Department,
		My_DTO.Leave_type,
	}

	err := My_Server.Base_DAO.Excused(Pending_Application)
	if err != nil {
		return err
	}

	return nil
}
