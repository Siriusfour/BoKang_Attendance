package DTO

type ViewDTO struct {
	UseID          int    `json:"UserID"         binding:"required"`
	AccessToken    string `json:"AccessToken"    binding:"required"`
	ApplicationsID int    `json:"ApplicationsID" binding:"required"`
	Pass           bool   `json:"Pass"           binding:"required"`
}
