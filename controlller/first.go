package controlller

import "github.com/gin-gonic/gin"

func First(c *gin.Context) {
	c.JSON(200, gin.H{
		"Code": "200",
		"Data": "hello this is our 117 web",
		"Msg":  "现在表示你可以正式访问到我们117web了，开始使用我们的API吧！",
	})
}
