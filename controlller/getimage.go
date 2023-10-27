package controlller

import (
	"117web/model/res"
	"117web/services"
	"github.com/gin-gonic/gin"
)

func GetImage(c *gin.Context) {
	imgname := c.Query("imagename")
	if file, err := services.GetImage(imgname); err != nil {
		res.FailWithCode(res.InternalServerError, c)
	} else {
		if _, err := c.Writer.WriteString(file); err != nil {
			res.FailWithCode(res.InternalServerError, c)
		}
	}
}
