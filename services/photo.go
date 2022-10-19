package services

import (
	"github.com/mrizalr/mygram/entities"
	"github.com/mrizalr/mygram/models"
	"github.com/mrizalr/mygram/repositories"
)

type PhotoService interface {
	Create(user entities.User, createRequest models.CreateRequest) (entities.Photo, error)
}

type photoService struct {
	repository *repositories.PhotoRepository
}

// func (s *service) Create(user entities.User, createRequest models.CreateRequest) (entities.Photo, error) {

// }
