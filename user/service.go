package user

import (
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Register(userRegisterRequest UserRegisterRequest) (User, error)
	Login(userLoginRequest UserLoginRequest) (User, error)
	Update(id int, userUpdateRequest UserUpdateRequest) (User, error)
	Delete(id int) (User, error)
}

type service struct {
	repository *repository
}

func NewService(repository *repository) *service {
	return &service{repository: repository}
}

func (s *service) Register(userRegisterRequest UserRegisterRequest) (User, error) {
	pwHash, err := bcrypt.GenerateFromPassword([]byte(userRegisterRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		return User{}, err
	}

	newUser := User{
		Username: userRegisterRequest.Username,
		Email:    userRegisterRequest.Email,
		Password: string(pwHash),
		Age:      userRegisterRequest.Age,
	}

	return s.repository.Create(newUser)
}

func (s *service) Login(userLoginRequest UserLoginRequest) (User, error) {
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

func (s *service) Update(id int, userUpdateRequest UserUpdateRequest) (User, error) {
	user, err := s.repository.FindByID(id)
	if err != nil {
		return user, err
	}

	user.Email = userUpdateRequest.Email
	user.Username = userUpdateRequest.Username

	return s.repository.Save(user)
}

func (s *service) Delete(id int) (User, error) {
	user, err := s.repository.FindByID(id)
	if err != nil {
		return user, err
	}

	return s.repository.Delete(user)
}
