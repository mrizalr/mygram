package user

import (
	"gorm.io/gorm"
)

type Repository interface {
	Create(user User) (User, error)
	FindByEmail(email string) (User, error)
	FindByID(id int) (User, error)
	Save(user User) (User, error)
	Delete(user User) (User, error)
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

func (r *repository) FindByEmail(email string) (User, error) {
	user := User{}
	err := r.db.Where("email = ?", email).First(&user).Error
	return user, err
}

func (r *repository) FindByID(id int) (User, error) {
	user := User{}
	err := r.db.Where("id = ?", id).First(&user).Error
	return user, err
}

func (r *repository) Save(user User) (User, error) {
	err := r.db.Save(user).Error
	return user, err
}

func (r *repository) Delete(user User) (User, error) {
	err := r.db.Delete(user).Error
	return user, err
}
