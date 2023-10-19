package model

type User struct {
	UserName string `json:"username",gorm:"size:255"`
	PassWord string `json:"password",gorm:"size:255"`
}
