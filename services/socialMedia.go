package services

import (
	"github.com/mrizalr/mygram/entities"
	"github.com/mrizalr/mygram/models"
	"github.com/mrizalr/mygram/repositories"
)

type SocmedService interface {
	Add(userID int, socmed models.AddSocialMediaRequest) (entities.SocialMedia, error)
}

type socmedService struct {
	socmedRepository repositories.SocmedRepository
}

func NewSocmedService(socmedRepository repositories.SocmedRepository) *socmedService {
	return &socmedService{
		socmedRepository: socmedRepository,
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
