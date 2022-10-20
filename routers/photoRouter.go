package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/mrizalr/mygram/handlers"
	"github.com/mrizalr/mygram/middlewares"
)

func InitPhotoRoutes(Routes *gin.Engine, handler *handlers.PhotoHandlers) {
	userGroup := Routes.Group("/photos")
	{
		userGroup.POST("/", middlewares.Auth, handler.UploadPhoto)
		userGroup.GET("/", middlewares.Auth, handler.GetAllPhotos)
	}
}
