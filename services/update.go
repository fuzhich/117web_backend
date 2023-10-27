package services

import "117web/dao"

func Update(uid string, filed string, val string) (bool, error) {
	if rs, err := dao.Update(uid, filed, val); err != nil {
		return false, err
	} else {
		if rs == true {
			return true, nil
		} else {
			return false, nil
		}
	}
}
