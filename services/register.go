package services

import (
	"117web/dao"
	"117web/model/res"
	"fmt"
)

func Register(un string, pw string) (res.Code, error) {
	if rs, err := dao.Register(un, pw); err != nil {
		fmt.Printf("%v", err)
		return res.InternalServerError, err
	} else {
		if rs == 0 {
			return res.DataMoreThanOne, nil
		} else {
			return res.OKCode, nil
		}
	}

}
