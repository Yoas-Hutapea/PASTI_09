package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Yoas-Hutapea/Microservice_09/auth/models"
	"github.com/Yoas-Hutapea/Microservice_09/auth/services"
)

type AuthHandler struct {
	AuthService *services.AuthService
}

func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	// Initialize and configure the AuthHandler instance
	return &AuthHandler{
		AuthService: authService,
		// Initialize other fields or dependencies
	}
}

func (ah *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	// Parse the request body
	var loginRequest models.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&loginRequest)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Call the login service
	token, user, err := ah.AuthService.Login(loginRequest.NIK, loginRequest.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// Create the response object
	loginResponse := models.LoginResponse{
		Token: token,
		User: models.UserDetail{
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
		},
	}

	// Return the response
	jsonResponse, err := json.Marshal(loginResponse)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

// Other authentication-related handler methods can be added here
