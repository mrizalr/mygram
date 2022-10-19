package database

import (
	"log"
	"os"

	"github.com/mrizalr/mygram/entities"
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

	db.AutoMigrate(&entities.User{}, &entities.Photo{}, &entities.Comment{}, &entities.SocialMedia{})
	log.Printf("Success connecting to database")
}

func GetDB() *gorm.DB {
	return db
}
