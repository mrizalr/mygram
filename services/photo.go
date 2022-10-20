package services

import (
	"errors"
	"strconv"

	"github.com/mrizalr/mygram/entities"
	"github.com/mrizalr/mygram/models"
	"github.com/mrizalr/mygram/repositories"
)

type PhotoService interface {
	Create(userID int, createRequest models.CreatePhotoRequest) (entities.Photo, error)
	GetAll() ([]entities.Photo, error)
	Update(paramID string, userID int, updateRequest models.CreatePhotoRequest) (entities.Photo, error)
	Delete(paramID string, userID int) (entities.Photo, error)
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

func (s *photoService) GetAll() ([]entities.Photo, error) {
	return s.photoRepository.Find()
}

func (s *photoService) Update(paramID string, userID int, updateRequest models.CreatePhotoRequest) (entities.Photo, error) {
	ID, err := strconv.Atoi(paramID)
	if err != nil {
		return entities.Photo{}, err
	}

	photo, err := s.photoRepository.FindByID(ID)
	if err != nil {
		return entities.Photo{}, err
	}

	if photo.UserID != uint(userID) {
		return entities.Photo{}, errors.New("Not authorized")
	}

	photo.Title = updateRequest.Title
	photo.Caption = updateRequest.Caption
	photo.PhotoURL = updateRequest.PhotoURL

	return s.photoRepository.Save(photo)
}

func (s *photoService) Delete(paramID string, userID int) (entities.Photo, error) {
	ID, err := strconv.Atoi(paramID)
	if err != nil {
		return entities.Photo{}, err
	}

	photo, err := s.photoRepository.FindByID(ID)
	if err != nil {
		return entities.Photo{}, err
	}

	if photo.UserID != uint(userID) {
		return entities.Photo{}, errors.New("Not authorized")
	}

	return s.photoRepository.Delete(photo)
}
