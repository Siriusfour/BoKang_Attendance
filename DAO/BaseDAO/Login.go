package BaseDAO

import "Attendance/Model"

func (My_BaseDAO *BaseDAO) Login(Name string, password string) Model.User {

	var My_User Model.User
	My_BaseDAO.orm.Model(&Model.User{}).Where("Name = ? AND password = ?", Name, password).First(&My_User)
	return My_User
}
