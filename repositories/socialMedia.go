package repositories

import (
	"gorm.io/gorm"
	"github.com/mrizalr/mygram/entities"
)

type SocmedRepository interface {
	Create(socmed entities.SocialMedia) (entities.SocialMedia, error)
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