package DTO

import (
	"gorm.io/gorm"
	"time"
)

// PageDTO 分页通用DTO
type PageDTO struct {
	PageIndex int `json:"pageIndex"`
	Limit     int `json:"limit"`
}

// GetPage 检验合法性
func (PageDTO *PageDTO) GetPage() {

	if PageDTO.PageIndex <= 0 {
		PageDTO.PageIndex = 1
	}
	if PageDTO.Limit <= 0 {
		PageDTO.Limit = 10
	}
}

// LoginDTO 从URL参数、请求头、请求体中拿到的信息，将来绑定到UserLoginDTO上
type LoginDTO struct {
	UserID        int    `json:"UserID" binding:"required" Message:"用户名错误！！" required_err:"用户名不能为空！！" `
	Refresh_Token string `json:"Refresh_Token"`
	Access_Token  string `json:"Access_Token"`
	Password      string `json:"Password"`
}

type ApplicationsDTO struct {
	gorm.Model
	Name       string    `json:"Name" binding:"required" `
	UserId     int       `json:"UserId" binding:"required" `
	Message    string    `json:"Message"`
	StartTime  time.Time `json:"StartTime" binding:"required"`
	EndTime    time.Time `json:"EndTime" binding:"required"`
	Department int       `json:"Department" binding:"required"`
	Leave_type int       `json:"Leave_type" binding:"required"`
	Token      string    `json:"Token" binding:"required"`
}

type ApplicationsDTOArray struct {
	MyApplications         []ApplicationsDTO
	DepartmentApplications []ApplicationsDTO
}

type UserInfoDTO struct {
	UserID       int    `gorm:"type:int;not null"`
	Name         string `gorm:"type:varchar(20);not null"`
	departmental string `gorm:"type:varchar(255);not null"`
	Leader       int    `gorm:"type:int;not null"`
}
