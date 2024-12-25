package LeaderDAO

import "gorm.io/gorm"
import "Attendance/Global"

type LeaderDAO struct {
	orm *gorm.DB
}

func New_Leader_DAO() *LeaderDAO {
	return &LeaderDAO{orm: Global.DB}
}
