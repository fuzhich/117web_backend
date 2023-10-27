package model

type Image struct {
	Name string `json:"imagename",gorm:"size:128"`
	Path string `json:"imagepath",gorm:"size:128"`
}
