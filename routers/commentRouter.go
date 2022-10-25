package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/mrizalr/mygram/handlers"
	"github.com/mrizalr/mygram/middlewares"
)

func InitCommentRoutes(Routes *gin.Engine, handler *handlers.CommentHandlers) {
	commentGroup := Routes.Group("/comments")
	{
		commentGroup.Use(middlewares.Auth)
		commentGroup.POST("/", handler.CreateComment)
		commentGroup.GET("/", handler.GetAllComment)
		commentGroup.PUT("/:commentId", handler.UpdateComment)
		commentGroup.DELETE("/:commentId", handler.DeleteComment)
	}
}
