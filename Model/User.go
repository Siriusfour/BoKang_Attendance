package Model

//定义要用到的表的结构

type User struct {
	UserID       int    `gorm:"type:int;not null"`
	Name         string `gorm:"type:varchar(20);not null"`
	Sex          int    `gorm:"type:INT;not null"`
	Password     string `gorm:"type:varchar(255);not null"`
	departmental string `gorm:"type:varchar(255);not null"`
	Leader       int    `gorm:"type:int;not null"`
}

func (User) TableName() string {
	return "users"
}
