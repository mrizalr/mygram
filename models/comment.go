package models

type CreateCommentRequest struct {
	Message string `json:"message" binding:"required"`
	PhotoId uint   `json:"photo_id" binding:"required,number"`
}
