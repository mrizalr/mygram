package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mrizalr/mygram/middlewares"
	"github.com/mrizalr/mygram/user"
)

func InitUserRoutes(Routes *gin.Engine, userHandler *user.UserHandler) {
	userGroup := Routes.Group("/users")
	{
		userGroup.POST("/register", userHandler.UserRegisterHandler)
		userGroup.POST("/login", userHandler.UserLoginHandler)
		userGroup.PUT("/", middlewares.Auth, userHandler.UserUpdateHandler)
		userGroup.DELETE("/", middlewares.Auth, userHandler.DeleteUserHandler)
	}
}
