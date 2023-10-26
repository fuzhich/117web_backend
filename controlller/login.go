package controlller

import (
	"117web/model/res"
	"117web/services"
	"fmt"
	"github.com/gin-gonic/gin"
)

type lg struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	//接收相应数据
	var recv lg
	if err := c.BindJSON(&recv); err != nil {
		fmt.Printf("接收数据失败")
		res.FailWithCode(res.BadRequest, c)
		return
	}
	//比较数据库中的账户和密码
	//返回结果
	if rs, err := services.Login(recv.Username, recv.Password); err != nil {
		res.FailWithCode(res.InternalServerError, c)
	} else {
		if rs == true {
			res.OKWithMessage("成功", c)
		} else {
			res.FailWithCode(res.NotMatch, c)
		}
	}
}
