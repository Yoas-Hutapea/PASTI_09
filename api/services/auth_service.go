package services

import (
	"errors"

	"github.com/Yoas-Hutapea/Microservice_09/api/models"
	"github.com/Yoas-Hutapea/Microservice_09/api/repositories"
)

type AuthService struct {
	UserRepository *repositories.UserRepository
}

func (as *AuthService) Login(nik, password string) error {
	// Implement the logic for user login using nik and password
	// Retrieve user by nik from the UserRepository
	// Compare the password with the stored password
	// Return an error if the login fails
	user, err := as.UserRepository.GetUserByNIK(nik)
	if err != nil {
		return err
	}

	if user.Password != password {
		return errors.New("Invalid password")
	}

	return nil
}