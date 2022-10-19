package services

import (
	"github.com/mrizalr/mygram/entities"
	"github.com/mrizalr/mygram/models"
	"github.com/mrizalr/mygram/repositories"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(userRegisterRequest models.UserRegisterRequest) (entities.User, error)
	Login(userLoginRequest models.UserLoginRequest) (entities.User, error)
	Update(id int, userUpdateRequest models.UserUpdateRequest) (entities.User, error)
	Delete(id int) (entities.User, error)
}

type userService struct {
	repository repositories.UserRepository
}

func NewUserService(repository repositories.UserRepository) *userService {
	return &userService{repository: repository}
}

func (s *userService) Register(userRegisterRequest models.UserRegisterRequest) (entities.User, error) {
	pwHash, err := bcrypt.GenerateFromPassword([]byte(userRegisterRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		return entities.User{}, err
	}

	newUser := entities.User{
		Username: userRegisterRequest.Username,
		Email:    userRegisterRequest.Email,
		Password: string(pwHash),
		Age:      userRegisterRequest.Age,
	}

	return s.repository.Create(newUser)
}

func (s *userService) Login(userLoginRequest models.UserLoginRequest) (entities.User, error) {
	userFound, err := s.repository.FindByEmail(userLoginRequest.Email)
	if err != nil {
		return userFound, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(userFound.Password), []byte(userLoginRequest.Password))
	if err != nil {
		return userFound, err
	}

	return userFound, nil
}

func (s *userService) Update(id int, userUpdateRequest models.UserUpdateRequest) (entities.User, error) {
	user, err := s.repository.FindByID(id)
	if err != nil {
		return user, err
	}

	user.Email = userUpdateRequest.Email
	user.Username = userUpdateRequest.Username

	return s.repository.Save(user)
}

func (s *userService) Delete(id int) (entities.User, error) {
	user, err := s.repository.FindByID(id)
	if err != nil {
		return user, err
	}

	return s.repository.Delete(user)
}
