package DTO

type ViewDTO struct {
	ApplicationsID int  `json:"ApplicationsID" binding:"required"`
	OK             bool `json:"OK"`
}
