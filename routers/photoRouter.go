package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/mrizalr/mygram/handlers"
	"github.com/mrizalr/mygram/middlewares"
)

func InitPhotoRoutes(Routes *gin.Engine, handler *handlers.PhotoHandlers) {
	photoGroup := Routes.Group("/photos")
	{
		photoGroup.Use(middlewares.Auth)
		photoGroup.POST("/", handler.UploadPhoto)
		photoGroup.GET("/", handler.GetAllPhotos)
		photoGroup.PUT("/:photoId", handler.UpdatePhoto)
		photoGroup.DELETE("/:photoId", handler.DeletePhoto)
	}
}
