package BaseDAO

import (
	"gorm.io/gorm"
)

type BaseDAO struct {
	orm *gorm.DB
}

func New_Base_DAO(orm *gorm.DB) *BaseDAO {
	return &BaseDAO{orm: orm}
}
