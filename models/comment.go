package models

import (
	"time"

	"github.com/mrizalr/mygram/entities"
)

type CreateCommentRequest struct {
	Message string `json:"message" binding:"required"`
	PhotoId uint   `json:"photo_id" binding:"required,number"`
}

type UpdateCommentRequest struct {
	Message string `json:"message"`
}

type GetCommentResponse struct {
	ID        uint      `json:"id"`
	UserID    uint      `json:"user_id"`
	PhotoID   uint      `json:"photo_id"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User      struct {
		ID       uint   `json:"id"`
		Email    string `json:"email"`
		Username string `json:"username"`
	}
}

func ParseToGetCommentResponse(comment entities.Comment, user entities.User) GetCommentResponse {
	return GetCommentResponse{
		ID:        comment.ID,
		UserID:    comment.UserID,
		PhotoID:   comment.PhotoID,
		Message:   comment.Message,
		CreatedAt: comment.CreatedAt,
		UpdatedAt: comment.UpdatedAt,
		User: struct {
			ID       uint   `json:"id"`
			Email    string `json:"email"`
			Username string `json:"username"`
		}{
			ID:       user.ID,
			Email:    user.Email,
			Username: user.Username,
		},
	}
}
