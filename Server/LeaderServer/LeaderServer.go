package LeaderServer

import (
	"Attendance/DAO/LeaderDAO"
)

type LeaderServer struct {
	LeaderDAO *LeaderDAO.LeaderDAO
}

func New_Leader_Server() *LeaderServer {
	return &LeaderServer{
		LeaderDAO: LeaderDAO.New_Leader_DAO(),
	}

}
