package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/mrizalr/mygram/handlers"
	"github.com/mrizalr/mygram/middlewares"
)

func InitSocmedRouter(Routes *gin.Engine, handler *handlers.SocialMediaHandlers) {
	socmedGroup := Routes.Group("/socialmedias")
	{
		socmedGroup.Use(middlewares.Auth)
		socmedGroup.POST("/", handler.AddSocialMedia)
		socmedGroup.GET("/", handler.GetAllSocmeds)
		socmedGroup.PUT("/:socialMediaId", handler.UpdateSocmed)
		socmedGroup.DELETE("/:socialMediaId", handler.DeleteSocmed)
	}
}
