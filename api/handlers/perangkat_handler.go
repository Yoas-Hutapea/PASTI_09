package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Yoas-Hutapea/Microservice_09/api/models"
	"github.com/Yoas-Hutapea/Microservice_09/api/services"
)

type PerangkatHandler struct {
	PerangkatService *services.PerangkatDesaService
}

func (ph *PerangkatHandler) AddPerangkat(w http.ResponseWriter, r *http.Request) {
	// Parse the request body
	var perangkat models.PerangkatDesa
	err := json.NewDecoder(r.Body).Decode(&perangkat)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Call the add perangkat service
	err = ph.PerangkatService.AddPerangkatDesa(&perangkat)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return success response
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, "Perangkat added successfully")
}

func (ph *PerangkatHandler) UpdatePerangkatDesa(w http.ResponseWriter, r *http.Request) {
	// Parse the request body
	var perangkat models.PerangkatDesa
	err := json.NewDecoder(r.Body).Decode(&perangkat)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Call the update perangkat service
	err = ph.PerangkatService.UpdatePerangkatDesa(&perangkat)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return success response
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Perangkat updated successfully")
}

func (ph *PerangkatHandler) DeletePerangkatDesa(w http.ResponseWriter, r *http.Request) {
	// Parse the request parameters
	// Assuming the perangkat ID is passed as a query parameter named "id"
	perangkatID := r.URL.Query().Get("id")
	if perangkatID == "" {
		http.Error(w, "Perangkat ID is required", http.StatusBadRequest)
		return
	}

	// Call the delete perangkat service
	err := ph.PerangkatService.DeletePerangkatDesa(perangkatID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return success response
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Perangkat deleted successfully")
}