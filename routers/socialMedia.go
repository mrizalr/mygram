package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/mrizalr/mygram/handlers"
	"github.com/mrizalr/mygram/middlewares"
)

func InitSocmedRouter(Routes *gin.Engine, handler *handlers.SocialMediaHandlers) {
	socmedGroup := Routes.Group("/socialmedias")
	{
		socmedGroup.POST("/", middlewares.Auth, handler.AddSocialMedia)
		socmedGroup.GET("/", middlewares.Auth, handler.GetAllSocmeds)
		socmedGroup.PUT("/:socialMediaId", middlewares.Auth, handler.UpdateSocmed)
	}
}
