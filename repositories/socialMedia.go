package repositories

import (
	"github.com/mrizalr/mygram/entities"
	"gorm.io/gorm"
)

type SocmedRepository interface {
	Create(socmed entities.SocialMedia) (entities.SocialMedia, error)
	FindAll() ([]entities.SocialMedia, error)
}

type socmedRepository struct {
	db *gorm.DB
}

func NewSocmedRepository(db *gorm.DB) *socmedRepository {
	return &socmedRepository{
		db: db,
	}
}

func (r *socmedRepository) Create(socmed entities.SocialMedia) (entities.SocialMedia, error) {
	return socmed, r.db.Create(&socmed).Error
}

func (r *socmedRepository) FindAll() ([]entities.SocialMedia, error){
	socialMedias := []entities.SocialMedia{}
	err := r.db.Find(&socialMedias).Error

	return socialMedias, err
}
