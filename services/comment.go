package services

import (
	"github.com/mrizalr/mygram/entities"
	"github.com/mrizalr/mygram/models"
	"github.com/mrizalr/mygram/repositories"
)

type CommentService interface {
	CreateComment(userID int, commentRequest models.CreateCommentRequest) (entities.Comment, error)
}

type commentService struct {
	commentRepository repositories.CommentRepository
}

func NewCommentService(commentRepository repositories.CommentRepository) *commentService {
	return &commentService{
		commentRepository: commentRepository,
	}
}

func (s *commentService) CreateComment(userID int, commentRequest models.CreateCommentRequest) (entities.Comment, error) {
	comment := entities.Comment{
		UserID:  uint(userID),
		PhotoID: commentRequest.PhotoId,
		Message: commentRequest.Message,
	}

	return s.commentRepository.Create(comment)
}
