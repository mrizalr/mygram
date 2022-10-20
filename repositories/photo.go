package repositories

import (
	"github.com/mrizalr/mygram/entities"

	"gorm.io/gorm"
)

type PhotoRepository interface {
	Create(user entities.User, photo entities.Photo) (entities.Photo, error)
	Find() ([]entities.Photo, error)
}

type photoRepository struct {
	db *gorm.DB
}

func NewPhotoRepository(db *gorm.DB) *photoRepository {
	return &photoRepository{
		db: db,
	}
}

func (r *photoRepository) Create(user entities.User, photo entities.Photo) (entities.Photo, error) {
	err := r.db.Model(&user).Association("Photos").Append(&photo)
	return photo, err
}

func (r *photoRepository) Find() ([]entities.Photo, error) {
	var photos []entities.Photo
	err := r.db.Model(&photos).Preload("User").Find(&photos).Error
	return photos, err
}
