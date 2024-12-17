package BaseServer

import (
	"Attendance/DAO/Base_DAO"
	"Attendance/Global"
)

type BaseServer struct {
	Base_DAO *Base_DAO.BaseDAO
}

func New_Base_Server() *BaseServer {

	DB := Base_DAO.NewBaseDAO(Global.DB)
	return &BaseServer{
		Base_DAO: &DB,
	}
}
