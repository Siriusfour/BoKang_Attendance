package DTO

type AddWorkDTO struct {
	ID              int
	Name            string `json:"Name"  binding:"required" Message:"用户名不能为空！"`
	PassWord        string `json:"Password"  binding:"required,omitempty" Message:"密码不能为空！"`
	ConfirmPassWord string `json:"confirm_password" binding:"required"`
	Email           string `json:"email" binding:"required"`
	Sex             int    `json:"gender" binding:"oneof=0 1 2"` // 性别 0:未知 1:男 2:女
}

type UserListDTO struct {
	PageDTO
}

type GetUserDTO struct {
}
