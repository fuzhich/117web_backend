package controlller

import (
	"117web/model/res"
	"117web/services"
	"fmt"
	"github.com/gin-gonic/gin"
)

type upd struct {
	Uid   string `json:"uid"`
	Feild string `json:"feild"`
	New   string `json:"new"`
}

func Update(c *gin.Context) {
	var recv upd
	if err := c.BindJSON(&recv); err != nil {
		fmt.Printf("接收数据失败")
		res.FailWithCode(res.BadRequest, c)
		return
	}
	if rs, err := services.Update(recv.Uid, recv.Feild, recv.New); err != nil {
		res.FailWithCode(res.InternalServerError, c)
	} else {
		if rs == true {
			res.OKWithMessage("成功", c)
		} else {
			res.FailWithCode(res.NotMatch, c)
		}
	}
}
