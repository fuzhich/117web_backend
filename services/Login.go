package services

import (
	"117web/dao"
)

func Login(un string, pw string) (bool, error) {
	//为空直接返回
	if un==""||pw==""{
		return false,nil
	}
	//比较密码
	if p, err := dao.Login(un); err != nil {
		return false, err
	} else {
		if p == pw {
			return true, nil
		} else {
			return false, nil
		}
	}
}
