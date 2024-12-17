package BaseServer

import (
	"Attendance/Controller/DTO"
)

func (My_Server *BaseServer) Excused(My_DTO DTO.ExcusedApplicationsDTO) error {

	err := My_Server.Base_DAO.Excused(My_DTO)
	if err != nil {
		return err
	}

	return nil
}
