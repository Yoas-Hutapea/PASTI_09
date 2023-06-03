package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Yoas-Hutapea/Microservice_09/api/models"
	"github.com/Yoas-Hutapea/Microservice_09/api/services"
)

type AuthHandler struct {
	AuthService *services.AuthService
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
	token, err := ah.AuthService.Login(loginRequest.NIK, loginRequest.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// Return the token in the response
	response := models.LoginResponse{Token: token}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

// Other authentication-related handler methods can be added here



