package model

type User struct {
	UserName string `json:"username",gorm:"size:255"`
	Password string `json:"password",gorm:"size:255"`
}
