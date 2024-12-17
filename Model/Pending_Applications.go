package Model

import (
	"gorm.io/gorm"
	"time"
)

type Pending_Application struct {
	gorm.Model
	Name       string `gorm:"type:varchar(20);not null"`
	UserID     uint   `gorm:"type:varchar(20);not null"`
	Message    string
	StartTime  time.Time `gorm:"type:datetime;not null"`
	EndTime    time.Time `gorm:"type:datetime;not null"`
	Department int       `json:"department" binding:"required"`
	Leave_type int       //0事假 1出差 2外派
}
