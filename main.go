package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/mrizalr/mygram/database"
	"github.com/mrizalr/mygram/handlers"
	_ "github.com/mrizalr/mygram/initializers"
	"github.com/mrizalr/mygram/repositories"
	"github.com/mrizalr/mygram/routers"
	"github.com/mrizalr/mygram/services"
)

func init() {
	database.Connect()
}

func main() {
	Routes := gin.Default()

	userRepository := repositories.NewUserRepository(database.GetDB())
	userService := services.NewUserService(userRepository)
	userHandler := handlers.NewUserHandler(userService)
	routers.InitUserRoutes(Routes, userHandler)

	photoRepository := repositories.NewPhotoRepository(database.GetDB())
	photoService := services.NewPhotoService(photoRepository, userRepository)
	photoHandler := handlers.NewPhotoHandlers(photoService)
	routers.InitPhotoRoutes(Routes, photoHandler)

	commentRepository := repositories.NewCommentRepository(database.GetDB())
	commentService := services.NewCommentService(commentRepository, userRepository)
	commentHandler := handlers.NewCommentHandlers(commentService)
	routers.InitCommentRoutes(Routes, commentHandler)

	socmedRepository := repositories.NewSocmedRepository(database.GetDB())
	socmedService := services.NewSocmedService(socmedRepository, userRepository)
	socmedHandler := handlers.NewSocialMediaHandlers(socmedService)
	routers.InitSocmedRouter(Routes, socmedHandler)

	Routes.Run(os.Getenv("SERVER_PORT"))
}
