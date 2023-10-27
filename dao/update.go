package dao

import (
	"117web/config"
	"117web/model"
	"fmt"
)

func Update(uid string, filed string, val string) (bool, error) {
	fmt.Println(uid + filed + val)
	if result := config.DB.Model(&model.Person{}).Where("uid = ?", uid).Update(filed, val); result.Error != nil {
		return false, result.Error
	}
	return true, nil
}
