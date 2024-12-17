package DAO

import (
	"Attendance/DAO/Base_DAO"
	"Attendance/Global"
)

type AdminDAO struct {
	Base_DAO.BaseDAO
}

func NewAdminDAO() {
	Base_DAO.NewBaseDAO(Global.DB)
}
