package Base_DAO

import (
	"Attendance/Controller/DTO"
)

func (My_BaseDAO *BaseDAO) Excused(My_DTO DTO.ExcusedApplicationsDTO) error {

	result := My_BaseDAO.orm.Create(My_DTO)

	if result.Error != nil {
		return result.Error
	}
	return nil

}
