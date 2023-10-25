package dao

import (
	"117web/config"
	"117web/model"
)

func Login(un string) (string, error) {
	var p []string
	//查询数据库该用户名的密码并返回
	if result := config.DB.Model(&model.User{}).Where("user_name = ?", un).Select("password").Find(&p); result.Error != nil {
		return "", result.Error
	} else {
		if len(p) == 0 {
			return "", nil
		} else {
			return p[0], nil
		}
	}

}
