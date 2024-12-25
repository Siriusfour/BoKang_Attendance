package LeaderServer

import "Attendance/Controller/DTO"

func (leader_server *LeaderServer) View(ViewDTO DTO.ViewDTO) error {

	err := leader_server.LeaderDAO.View(ViewDTO.ApplicationsID, ViewDTO.OK)
	if err != nil {
		return err
	}

	return nil
}
