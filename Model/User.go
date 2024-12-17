package Model

//定义要用到的表的结构

import (
	"Attendance/Utills"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserID       int    `gorm:"type:int;not null"`
	Name         string `gorm:"type:varchar(20);not null"`
	Sex          int    `gorm:"type:INT;not null"`
	Password     string `gorm:"type:varchar(255);not null"`
	departmental string `gorm:"type:varchar(255);not null"`
}

func EnCodePaasWord(password string) (string, error) {
	hash, err := Utills.EnCodePassWord(password)
	if err != nil {
		return "", err
	}
	return hash, err
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {

	user.Password, err = EnCodePaasWord(user.Password)

	return err
}
