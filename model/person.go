package model

type Person struct {
	RealName string `json:"realname",grom:"size:64"`
	Uid      int    `json:"uid",grom:"uid"`
	Gender   string `json:"gender",gorm:"size:8"`
}
