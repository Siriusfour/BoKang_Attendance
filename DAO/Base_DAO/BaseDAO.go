package Base_DAO

import (
	"Attendance/Model"
	"gorm.io/gorm"
)

type BaseDAO struct {
	orm *gorm.DB
}

func NewBaseDAO(orm *gorm.DB) BaseDAO {
	DB := BaseDAO{orm: orm}
	return DB
}

func (My_BaseDAO *BaseDAO) Get_User_Name_PassWord(UserID int, Password string) Model.User {

	var iuser Model.User
	My_BaseDAO.orm.Model(&iuser).Where("user_id = ? AND Password = ?", UserID, Password).Find(&iuser)
	return iuser

}
