package services

import (
	"errors"

	"github.com/Yoas-Hutapea/Microservice_09/api/repositories"
	"github.com/dgrijalva/jwt-go"
)

type AuthService struct {
	UserRepository *repositories.UserRepository
}

func NewAuthService(userRepository *repositories.UserRepository) *AuthService {
	return &AuthService{
		UserRepository: userRepository,
	}
}

func (as *AuthService) Login(nik, password string) (string, error) {
	user, err := as.UserRepository.GetUserByNIK(nik)
	if err != nil {
		return "", err
	}

	if user.Password != password {
		return "", errors.New("Invalid password")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"nik": user.NIK,
		// Add more claims as needed
	})

	tokenString, err := token.SignedString([]byte("orhutapea123"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
