package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"

	"github.com/Yoas-Hutapea/Microservice_09/penduduk/models"
	"github.com/Yoas-Hutapea/Microservice_09/penduduk/services"

)

type PendudukHandler struct {
	PendudukService *services.PendudukService
}

func NewPendudukHandler(pendudukService *services.PendudukService) *PendudukHandler {
	// Initialize and configure the PendudukHandler instance
	return &PendudukHandler{
		PendudukService: pendudukService,
		// Initialize other fields or dependencies
	}
}

func (ph *PendudukHandler) AddUser(w http.ResponseWriter, r *http.Request) {
	// Parse the request body
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Call the add user service
	err = ph.PendudukService.AddUser(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return success response
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, "User added successfully")
}

func (ph *PendudukHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	// Parse the request body
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Call the update user service
	err = ph.PendudukService.UpdateUser(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return success response
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "User updated successfully")
}

func (ph *PendudukHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	// Parse the request parameters
	// Assuming the user ID is passed as a path variable named "id"
	vars := mux.Vars(r)
	userID := vars["id"]
	if userID == "" {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}

	// Convert userID to int
	id, err := strconv.Atoi(userID)
	if err != nil {
		http.Error(w, "Invalid User ID", http.StatusBadRequest)
		return
	}

	// Call the delete user service
	err = ph.PendudukService.DeleteUser(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return success response
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "User deleted successfully")
}

