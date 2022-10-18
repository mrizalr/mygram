package user

import "gorm.io/gorm"

type Repository interface {
	Create(user User) (User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Create(user User) (User, error) {
	err := r.db.Create(&user).Error
	return user, err
}