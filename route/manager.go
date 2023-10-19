package route

import (
	"117web/controlller"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.New()
	sso := router.Group("/sso")
	{
		sso.GET("", controlller.First)
	}
	return router
}
