package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/mrizalr/mygram/handlers"
	"github.com/mrizalr/mygram/middlewares"
)

func InitUserRoutes(Routes *gin.Engine, handler *handlers.UserHandler) {
	userGroup := Routes.Group("/users")
	{
		userGroup.POST("/register", handler.UserRegisterHandler)
		userGroup.POST("/login", handler.UserLoginHandler)

		userGroup.Use(middlewares.Auth)
		userGroup.PUT("/", handler.UserUpdateHandler)
		userGroup.DELETE("/", handler.DeleteUserHandler)
	}
}
