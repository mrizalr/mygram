package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mrizalr/mygram/user"
)

func InitUserRoutes(Routes *gin.Engine, userHandler user.Handler) {
	userGroup := Routes.Group("/users")
	{
		userGroup.POST("/register", userHandler.UserRegisterHandler)
	}
}
