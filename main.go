package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/mrizalr/mygram/database"
	_ "github.com/mrizalr/mygram/initializers"
	"github.com/mrizalr/mygram/router"
	"github.com/mrizalr/mygram/user"
)

func init() {
	database.Connect()
}

func main() {
	Routes := gin.Default()

	userRepository := user.NewRepository(database.GetDB())
	userService := user.NewService(userRepository)
	userHandler := user.NewHandler(userService)
	router.InitUserRoutes(Routes, userHandler)

	Routes.Run(os.Getenv("SERVER_PORT"))
}
