package Model

import (
	"gorm.io/gorm"
	"time"
)

type Application struct {
	gorm.Model
	Name       string    `gorm:"type:varchar(20);not null"`
	UserID     int       `gorm:"type:int;not null"`
	Message    string    `gorm:"type:varchar(255);not null"`
	StartTime  time.Time `gorm:"type:datetime;not null"`
	EndTime    time.Time `gorm:"type:datetime;not null"`
	Department int       `gorm:"type:varchar(20);not null"`
	Leave_type int       `gorm:"type:varchar(20);not null"` //0事假 1出差 2外派
	Status     int       `gorm:"type:int;not null"`
}

func (Application) TableName() string {
	return "applications" // 指定表名为 "applications"
}
