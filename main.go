package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/mrizalr/mygram/database"
	_ "github.com/mrizalr/mygram/initializers"
)

var Routes *gin.Engine

func init() {
	database.Connect()
}

func main() {
	Routes = gin.Default()
	Routes.Run(os.Getenv("SERVER_PORT"))
}
