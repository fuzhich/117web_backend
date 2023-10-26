package controlller

import (
	"117web/model/res"
	"117web/services"
	"fmt"
	"github.com/gin-gonic/gin"
)

type rg struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Register(c *gin.Context) {
	var recv rg
	if err := c.BindJSON(&recv); err != nil {
		fmt.Printf("接收数据失败")
		res.FailWithCode(res.BadRequest, c)
		return
	}
	if code, err := services.Register(recv.Username, recv.Password); err != nil {
		res.FailWithCode(res.InternalServerError, nil)
	} else {
		res.OKWithCode(code, c)
	}
}
