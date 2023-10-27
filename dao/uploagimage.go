package dao

import (
	"117web/config"
	"117web/model"
)

func UploadImage(path string, name string) error {
	var rs []string
	if result := config.DB.Model(&model.Image{}).Where("name = ?", name).Select("path").Find(&rs); result.Error != nil {
		return result.Error
	}
	if len(rs) == 0 {
		if result := config.DB.Create(&model.Image{Name: name, Path: path}); result.Error != nil {
			return result.Error
		}
	}

	return nil

}
