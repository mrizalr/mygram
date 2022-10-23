package repositories

import (
	"github.com/mrizalr/mygram/entities"
	"gorm.io/gorm"
)

type CommentRepository interface {
	Create(comment entities.Comment) (entities.Comment, error)
	FindAll() ([]entities.Comment, error)
	FindByID(ID int) (entities.Comment, error)
	Save(comment entities.Comment) (entities.Comment, error)
	Delete(comment entities.Comment) (entities.Comment, error)
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

func (r *commentRepository) FindAll() ([]entities.Comment, error) {
	var comments []entities.Comment
	err := r.db.Find(&comments).Error
	return comments, err
}

func (r *commentRepository) FindByID(ID int) (entities.Comment, error) {
	var comment entities.Comment
	err := r.db.Where("id = ?", ID).First(&comment).Error
	return comment, err
}

func (r *commentRepository) Save(comment entities.Comment) (entities.Comment, error) {
	return comment, r.db.Save(&comment).Error
}

func (r *commentRepository) Delete(comment entities.Comment) (entities.Comment, error) {
	return comment, r.db.Delete(&comment).Error
}
