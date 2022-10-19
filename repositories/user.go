package repositories

import (
	"github.com/mrizalr/mygram/entities"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user entities.User) (entities.User, error)
	FindByEmail(email string) (entities.User, error)
	FindByID(id int) (entities.User, error)
	Save(user entities.User) (entities.User, error)
	Delete(user entities.User) (entities.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) Create(user entities.User) (entities.User, error) {
	err := r.db.Create(&user).Error
	return user, err
}

func (r *userRepository) FindByEmail(email string) (entities.User, error) {
	user := entities.User{}
	err := r.db.Where("email = ?", email).First(&user).Error
	return user, err
}

func (r *userRepository) FindByID(id int) (entities.User, error) {
	user := entities.User{}
	err := r.db.Where("id = ?", id).First(&user).Error
	return user, err
}

func (r *userRepository) Save(user entities.User) (entities.User, error) {
	err := r.db.Save(&user).Error
	return user, err
}

func (r *userRepository) Delete(user entities.User) (entities.User, error) {
	err := r.db.Delete(&user).Error
	return user, err
}
