package repositories

import (
	"github.com/mrizalr/mygram/entities"
	"gorm.io/gorm"
)

type SocmedRepository interface {
	Create(socmed entities.SocialMedia) (entities.SocialMedia, error)
	FindAll() ([]entities.SocialMedia, error)
	Save(socmed entities.SocialMedia) (entities.SocialMedia, error)
	FindByID(ID int) (entities.SocialMedia, error)
	Delete(socmed entities.SocialMedia) (entities.SocialMedia, error)
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

func (r *socmedRepository) FindAll() ([]entities.SocialMedia, error) {
	socialMedias := []entities.SocialMedia{}
	err := r.db.Find(&socialMedias).Error

	return socialMedias, err
}

func (r *socmedRepository) Save(socmed entities.SocialMedia) (entities.SocialMedia, error) {
	return socmed, r.db.Save(&socmed).Error
}

func (r *socmedRepository) FindByID(ID int) (entities.SocialMedia, error) {
	socmed := entities.SocialMedia{}
	err := r.db.Where("id = ?", ID).First(&socmed).Error
	return socmed, err
}

func (r *socmedRepository) Delete(socmed entities.SocialMedia) (entities.SocialMedia, error) {
	return socmed, r.db.Delete(&socmed).Error
}
