package services

import (
	"github.com/mrizalr/mygram/entities"
	"github.com/mrizalr/mygram/models"
	"github.com/mrizalr/mygram/repositories"
)

type PhotoService interface {
	Create(userID int, createRequest models.CreatePhotoRequest) (entities.Photo, error)
}

type photoService struct {
	photoRepository repositories.PhotoRepository
	userRepository  repositories.UserRepository
}

func NewPhotoService(photoRepository repositories.PhotoRepository, userRepository repositories.UserRepository) *photoService {
	return &photoService{
		photoRepository: photoRepository,
		userRepository:  userRepository,
	}
}

func (s *photoService) Create(userID int, createRequest models.CreatePhotoRequest) (entities.Photo, error) {
	newPhoto := entities.Photo{
		Title:    createRequest.Title,
		Caption:  createRequest.Caption,
		PhotoURL: createRequest.PhotoURL,
	}

	user, err := s.userRepository.FindByID(userID)
	if err != nil {
		return newPhoto, err
	}

	return s.photoRepository.Create(user, newPhoto)
}
