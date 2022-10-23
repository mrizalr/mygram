package services

import (
	"errors"

	"github.com/mrizalr/mygram/entities"
	"github.com/mrizalr/mygram/models"
	"github.com/mrizalr/mygram/repositories"
)

type CommentService interface {
	Create(userID int, commentRequest models.CreateCommentRequest) (entities.Comment, error)
	GetAll() ([]models.GetCommentResponse, error)
	Update(commentID int, userID int, commentRequest models.UpdateCommentRequest) (entities.Comment, error)
	Delete(commentID int, userID int) (entities.Comment, error)
}

type commentService struct {
	commentRepository repositories.CommentRepository
	userRepository    repositories.UserRepository
}

func NewCommentService(commentRepository repositories.CommentRepository, userRepository repositories.UserRepository) *commentService {
	return &commentService{
		commentRepository: commentRepository,
		userRepository:    userRepository,
	}
}

func (s *commentService) Create(userID int, commentRequest models.CreateCommentRequest) (entities.Comment, error) {
	comment := entities.Comment{
		UserID:  uint(userID),
		PhotoID: commentRequest.PhotoId,
		Message: commentRequest.Message,
	}

	return s.commentRepository.Create(comment)
}

func (s *commentService) GetAll() ([]models.GetCommentResponse, error) {
	var Response []models.GetCommentResponse
	comments, err := s.commentRepository.FindAll()
	if err != nil {
		return Response, err
	}

	for _, comment := range comments {
		user, err := s.userRepository.FindByID(int(comment.UserID))
		if err != nil {
			return Response, err
		}

		c := models.ParseToGetCommentResponse(comment, user)
		Response = append(Response, c)
	}

	return Response, nil
}

func (s *commentService) Update(commentID int, userID int, commentRequest models.UpdateCommentRequest) (entities.Comment, error) {
	comment, err := s.commentRepository.FindByID(commentID)
	if err != nil {
		return comment, err
	}

	if comment.UserID != uint(userID) {
		return comment, errors.New("Unauthorized")
	}

	comment.Message = commentRequest.Message
	return s.commentRepository.Save(comment)
}

func (s *commentService) Delete(commentID int, userID int) (entities.Comment, error) {
	comment, err := s.commentRepository.FindByID(commentID)
	if err != nil {
		return comment, err
	}

	if comment.UserID != uint(userID) {
		return comment, errors.New("Unauthorized")
	}

	return s.commentRepository.Delete(comment)
}
