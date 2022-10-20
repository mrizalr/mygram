package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/mrizalr/mygram/handlers"
	"github.com/mrizalr/mygram/middlewares"
)

func InitCommentRoutes(Routes *gin.Engine, handler *handlers.CommentHandlers) {
	commentGroup := Routes.Group("/comments")
	{
		commentGroup.POST("/", middlewares.Auth, handler.CreateComment)
	}
}
