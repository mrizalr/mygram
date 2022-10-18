package database

import (
	"log"
	"os"

	"github.com/mrizalr/mygram/comment"
	"github.com/mrizalr/mygram/photo"
	socialmedia "github.com/mrizalr/mygram/socialMedia"
	"github.com/mrizalr/mygram/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host = "localhost"
	port = "5432"
)

var (
	db  *gorm.DB
	err error
)

func Connect() {
	db, err = gorm.Open(postgres.Open(os.Getenv("DB_CONFIG")), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting database")
	}

	db.AutoMigrate(&user.User{}, &photo.Photo{}, &comment.Comment{}, &socialmedia.SocialMedia{})
	log.Printf("Success connecting to database")
}

func GetDB() *gorm.DB {
	return db
}
