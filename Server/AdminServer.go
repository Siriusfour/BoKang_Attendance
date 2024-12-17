package Server

import (
	"Attendance/Controller/DTO"
	"Attendance/DAO"
)

type AdminServir struct {
	DAO *DAO.AdminDAO
}

func (adminservir *AdminServir) Admin_AddWork(Add_Work_DTO *DTO.AddWorkDTO) {

}
