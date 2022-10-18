package user

import (
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Register(userRequest UserRequest) (User, error)
}

type service struct {
	repository *repository
}

func NewService(repository *repository) *service {
	return &service{repository: repository}
}

func (s *service) Register(userRequest UserRequest) (User, error) {
	pwHash, err := bcrypt.GenerateFromPassword([]byte(userRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		return User{}, err
	}

	newUser := User{
		Username: userRequest.Username,
		Email:    userRequest.Email,
		Password: string(pwHash),
		Age:      userRequest.Age,
	}

	return s.repository.Create(newUser)
}
