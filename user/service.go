package user

import (
	"os"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Register(userRegisterRequest UserRegisterRequest) (User, error)
	Login(userLoginRequest UserLoginRequest) (LoginResponse, error)
	Update(id int, userUpdateRequest UserUpdateRequest) (User, error)
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

func (s *service) Login(userLoginRequest UserLoginRequest) (LoginResponse, error) {
	loginResponse := LoginResponse{}
	userFound, err := s.repository.FindByEmail(userLoginRequest.Email)
	if err != nil {
		return loginResponse, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(userFound.Password), []byte(userLoginRequest.Password))
	if err != nil {
		return loginResponse, err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"id": userFound.ID})
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return loginResponse, err
	}

	loginResponse.Token = tokenString
	return loginResponse, nil
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
