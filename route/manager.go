package route

import (
	"117web/controlller"
	"117web/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.New()
	sso := router.Group("/sso")
	{
		sso.GET("", middleware.Cors(), controlller.First)
		sso.POST("/login", middleware.Cors(), controlller.Login)
		sso.POST("/register", middleware.Cors(), controlller.Register)
	}
	return router
}
