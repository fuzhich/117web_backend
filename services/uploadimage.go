package services

import (
	"117web/dao"
	"117web/model/res"
	"fmt"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"path"
)

func UploadImage(file *multipart.FileHeader, c *gin.Context) error {
	filepath := path.Join("D:\\117webimages\\" + file.Filename)
	err := c.SaveUploadedFile(file, filepath)

	if err != nil {
		fmt.Printf("文件保存在路径：%s下失败", filepath)
		res.FailWithMessage("保存文件失败", c)
		return err
	}
	if err := dao.UploadImage(filepath, file.Filename); err != nil {
		return err
	}
	fmt.Printf("\n文件传入，保存在：%s", filepath)
	return nil
}
