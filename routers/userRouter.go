package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mrizalr/mygram/handlers"
	"github.com/mrizalr/mygram/middlewares"
)

func InitUserRoutes(Routes *gin.Engine, userHandler *handlers.UserHandler) {
	userGroup := Routes.Group("/users")
	{
		userGroup.POST("/register", userHandler.UserRegisterHandler)
		userGroup.POST("/login", userHandler.UserLoginHandler)
		userGroup.PUT("/", middlewares.Auth, userHandler.UserUpdateHandler)
		userGroup.DELETE("/", middlewares.Auth, userHandler.DeleteUserHandler)
	}
}
