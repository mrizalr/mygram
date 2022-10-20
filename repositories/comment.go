package repositories

import (
	"github.com/mrizalr/mygram/entities"
	"gorm.io/gorm"
)

type CommentRepository interface {
	Create(comment entities.Comment) (entities.Comment, error)
}

type commentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) *commentRepository {
	return &commentRepository{
		db: db,
	}
}

func (r *commentRepository) Create(comment entities.Comment) (entities.Comment, error) {
	return comment, r.db.Create(&comment).Error
}
