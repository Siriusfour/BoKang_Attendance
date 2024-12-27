package LeaderServer

import (
	"Attendance/Controller/DTO"
	"Attendance/Utills"
)

func (leader_server *LeaderServer) View(ViewDTO *DTO.ViewDTO) *Utills.MyError {

	err := leader_server.LeaderDAO.View(ViewDTO)
	if err != nil {
		return err
	}

	return nil
}
