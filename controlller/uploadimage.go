package controlller

import (
	"117web/model/res"
	"117web/services"
	"github.com/gin-gonic/gin"
)

func UploadImage(c *gin.Context) {
	if file, err := c.FormFile("image"); err != nil {
		res.FailWithCode(res.BadRequest, c)
		return
	} else {
		if err := services.UploadImage(file, c); err != nil {
			res.FailWithCode(res.InternalServerError, c)
		}
		res.OKWithCode(res.OKCode, c)
		return
	}
}
