package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/mrizalr/mygram/handlers"
	"github.com/mrizalr/mygram/middlewares"
)

func InitPhotoRoutes(Routes *gin.Engine, handler *handlers.PhotoHandlers) {
	photoGroup := Routes.Group("/photos")
	{
		photoGroup.POST("/", middlewares.Auth, handler.UploadPhoto)
		photoGroup.GET("/", middlewares.Auth, handler.GetAllPhotos)
		photoGroup.PUT("/:photoId", middlewares.Auth, handler.UpdatePhoto)
	}
}
