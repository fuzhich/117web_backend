package dao

import (
	"117web/config"
	"117web/model"
)

func Register(un string, pw string) (int, error) {
	var a []string
	//如果数据库存在该用户名则返回-1
	if result := config.DB.Model(&model.User{}).Where("user_name = ?", un).Select("user_name").Find(&a); result.Error != nil {
		return -1, result.Error
	}
	if len(a) > 0 {
		return 0, nil
	}
	//数据库中不存在该用户名数据，创建数据
	if result := config.DB.Create(&model.User{UserName: un, Password: pw}); result.Error != nil {
		return -1, result.Error
	} else {
		return 1, nil
	}

}
