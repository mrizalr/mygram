package services

import (
	"github.com/mrizalr/mygram/entities"
	"github.com/mrizalr/mygram/models"
	"github.com/mrizalr/mygram/repositories"
)

type SocmedService interface {
	Add(userID int, socmed models.AddSocialMediaRequest) (entities.SocialMedia, error)
	GetAll() ([]models.GetSocmedResponse, error)
}

type socmedService struct {
	socmedRepository repositories.SocmedRepository
	userRepository   repositories.UserRepository
}

func NewSocmedService(socmedRepository repositories.SocmedRepository, userRepository repositories.UserRepository) *socmedService {
	return &socmedService{
		socmedRepository: socmedRepository,
		userRepository:   userRepository,
	}
}

func (s *socmedService) Add(userID int, socmedRequest models.AddSocialMediaRequest) (entities.SocialMedia, error) {
	socmed := entities.SocialMedia{
		Name:           socmedRequest.Name,
		SocialMediaURL: socmedRequest.SocialMediaURL,
		UserID:         uint(userID),
	}

	return s.socmedRepository.Create(socmed)
}

func (s *socmedService) GetAll() ([]models.GetSocmedResponse, error) {
	response := []models.GetSocmedResponse{}
	socmeds, err := s.socmedRepository.FindAll()
	if err != nil {
		return response, err
	}

	for _, socmed := range socmeds {
		user, err := s.userRepository.FindByID(int(socmed.UserID))
		if err != nil {
			return response, err
		}

		response = append(response, models.ParseToGetSocmedResponse(socmed, user))
	}

	return response, nil
}
