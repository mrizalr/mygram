package repositories

import (
	"github.com/mrizalr/mygram/entities"

	"gorm.io/gorm"
)

type PhotoRepository interface {
	Create(user entities.User, photo entities.Photo) (entities.Photo, error)
	Find() ([]entities.Photo, error)
	FindByID(ID int) (entities.Photo, error)
	Save(photo entities.Photo) (entities.Photo, error)
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

func (r *photoRepository) FindByID(ID int) (entities.Photo, error) {
	var photo entities.Photo
	err := r.db.Where("id = ?", ID).Find(&photo).Error
	return photo, err
}

func (r *photoRepository) Save(photo entities.Photo) (entities.Photo, error) {
	return photo, r.db.Save(&photo).Error
}
