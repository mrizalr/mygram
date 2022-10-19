package repositories

import (
	"github.com/mrizalr/mygram/entities"

	"gorm.io/gorm"
)

type PhotoRepository interface {
	Create(user entities.User, photo entities.Photo) (entities.Photo, error)
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
