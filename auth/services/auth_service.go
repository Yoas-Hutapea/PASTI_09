package services

import (
	"errors"

	"golang.org/x/crypto/bcrypt"

	"github.com/Yoas-Hutapea/Microservice_09/auth/models"
	"github.com/Yoas-Hutapea/Microservice_09/auth/repositories"
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

func (as *AuthService) Login(nik, password string) (string, models.UserDetail, error) {
	// Check if the provided credentials are valid
	user, err := as.UserRepository.GetUserByNIK(nik)
	if err != nil {
		return "", models.UserDetail{}, err
	}

	// Compare the password with the hashed password stored in the user record
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", models.UserDetail{}, errors.New("invalid password")
	}

	// Generate a JWT token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["userID"] = user.ID
	// Add other claims as needed
	// e.g., claims["role"] = user.Role

	// Sign the token with a secret key
	tokenString, err := token.SignedString([]byte("your-secret-key"))
	if err != nil {
		return "", models.UserDetail{}, err
	}

	// Convert *models.User to models.UserDetail
	userDetail := models.UserDetail{
		ID:           user.ID,
		Nama:         user.Nama,
		NIK:          user.NIK,
		NoTelp:       user.NoTelp,
		Alamat:       user.Alamat,
		TempatLahir:  user.TempatLahir,
		TanggalLahir: user.TanggalLahir,
		Usia:         user.Usia,
		JenisKelamin: user.JenisKelamin,
		Pekerjaan:    user.Pekerjaan,
		Agama:        user.Agama,
		KK:           user.KK,
		Gambar:       user.Gambar,
	}

	return tokenString, userDetail, nil
}

func (as *AuthService) GetUserByNIK(userNIK string) (*models.User, error) {
	return as.UserRepository.GetUserByNIK(userNIK)
}

func (as *AuthService) GetAllUsers() ([]*models.User, error) {
	return as.UserRepository.GetAllUsers()
}