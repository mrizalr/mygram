package models

import (
	"time"

	"github.com/mrizalr/mygram/entities"
)

type AddSocialMediaRequest struct {
	Name           string `json:"name" binding:"required"`
	SocialMediaURL string `json:"social_media_url" binding:"required"`
}

type GetSocmedResponse struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	SocialMediaURL string    `json:"social_media_url"`
	UserID         int       `json:"user_id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	User           struct {
		ID       int    `json:"id"`
		Username string `json:"username"`
		Email    string `json:"email"`
	}
}

func ParseToGetSocmedResponse(socmed entities.SocialMedia, user entities.User) GetSocmedResponse {
	return GetSocmedResponse{
		ID:             int(socmed.ID),
		Name:           socmed.Name,
		SocialMediaURL: socmed.SocialMediaURL,
		UserID:         int(socmed.UserID),
		CreatedAt:      socmed.CreatedAt,
		UpdatedAt:      socmed.UpdatedAt,
		User: struct {
			ID       int    `json:"id"`
			Username string `json:"username"`
			Email    string `json:"email"`
		}{
			ID:       int(user.ID),
			Username: user.Username,
			Email:    user.Email,
		},
	}
}
