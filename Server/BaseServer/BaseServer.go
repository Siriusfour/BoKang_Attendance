package BaseServer

import (
	"Attendance/DAO/BaseDAO"
	"Attendance/Global"
)

type BaseServer struct {
	BaseDAO *BaseDAO.BaseDAO
}

func New_Base_Server() *BaseServer {

	DB := BaseDAO.New_Base_DAO(Global.DB)
	return &BaseServer{
		BaseDAO: DB,
	}
}
