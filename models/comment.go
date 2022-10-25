package models

import (
	"time"

	"github.com/mrizalr/mygram/entities"
)

type CreateCommentRequest struct {
	Message string `json:"message" binding:"required"`
	PhotoId uint   `json:"photo_id" binding:"required,number"`
}

type CreateCommentResponse struct {
	ID        uint      `json:"id"`
	Message   string    `json:"message"`
	PhotoID   uint      `json:"photo_id"`
	UserID    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

func ParseToCreateCommentResponse(comment entities.Comment) CreateCommentResponse {
	return CreateCommentResponse{
		ID:        comment.ID,
		Message:   comment.Message,
		PhotoID:   comment.PhotoID,
		UserID:    comment.UserID,
		CreatedAt: comment.CreatedAt,
	}
}

type UpdateCommentRequest struct {
	Message string `json:"message"`
}

type UpdateCommentResponse struct {
	ID        uint      `json:"id"`
	Message   string    `json:"message"`
	PhotoID   uint      `json:"photo_id"`
	UserID    uint      `json:"user_id"`
	UpdatedAt time.Time `json:"updated_at"`
}

func ParseToUpdateCommentResponse(comment entities.Comment) UpdateCommentResponse {
	return UpdateCommentResponse{
		ID:        comment.ID,
		Message:   comment.Message,
		PhotoID:   comment.PhotoID,
		UserID:    comment.UserID,
		UpdatedAt: comment.UpdatedAt,
	}
}

type GetCommentResponse struct {
	ID        uint      `json:"id"`
	Message   string    `json:"message"`
	PhotoID   uint      `json:"photo_id"`
	UserID    uint      `json:"user_id"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
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
