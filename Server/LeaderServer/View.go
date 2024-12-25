package LeaderServer

import "Attendance/Controller/DTO"

func (leader_server *LeaderServer) View(ViewDTO DTO.ViewDTO) {

	if ViewDTO.OK {
		leader_server.LeaderDAO.View_Access(ViewDTO.ApplicationsID)
	} else {
		leader_server.LeaderDAO.View_rejected(ViewDTO.ApplicationsID)
	}

}
