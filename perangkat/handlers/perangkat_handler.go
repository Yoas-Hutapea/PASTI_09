package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"

	"github.com/Yoas-Hutapea/Microservice_09/perangkat/models"
	"github.com/Yoas-Hutapea/Microservice_09/perangkat/services"
)

type PerangkatHandler struct {
	PerangkatService *services.PerangkatDesaService
}

func NewPerangkatHandler(perangkatService *services.PerangkatDesaService) *PerangkatHandler {
	// Initialize and configure the PendudukHandler instance
	return &PerangkatHandler{
		PerangkatService: perangkatService,
		// Initialize other fields or dependencies
	}
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
	params := mux.Vars(r)
	perangkatID := params["id"]
	if perangkatID == "" {
		http.Error(w, "Perangkat ID is required", http.StatusBadRequest)
		return
	}

	// Convert perangkatID to int
	id, err := strconv.Atoi(perangkatID)
	if err != nil {
		http.Error(w, "Invalid Perangkat ID", http.StatusBadRequest)
		return
	}

	// Call the delete perangkat service
	err = ph.PerangkatService.DeletePerangkatDesa(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return success response
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Perangkat deleted successfully")
}


func (ph *PerangkatHandler) GetPerangkatDesaByID(w http.ResponseWriter, r *http.Request) {
	// Parse the request parameters
	// Assuming the perangkat ID is passed as a path variable named "id"
	vars := mux.Vars(r)
	perangkatID := vars["id"]
	if perangkatID == "" {
		http.Error(w, "PerangkatDesa ID is required", http.StatusBadRequest)
		return
	}

	// Convert perangkatID to int
	id, err := strconv.Atoi(perangkatID)
	if err != nil {
		http.Error(w, "Invalid PerangkatDesa ID", http.StatusBadRequest)
		return
	}
// Call the get perangkat by ID service
perangkat, err := ph.PerangkatService.GetPerangkatDesaByID(id)
if err != nil {
	http.Error(w, "Error retrieving perangkat", http.StatusInternalServerError)
	return
}

// Marshal perangkat to JSON
jsonResponse, err := json.Marshal(perangkat)
if err != nil {
	http.Error(w, "Internal server error", http.StatusInternalServerError)
	return
}

// Set response headers and write JSON response
w.Header().Set("Content-Type", "application/json")
w.WriteHeader(http.StatusOK)
w.Write(jsonResponse)
}


func (ph *PerangkatHandler) GetAllPerangkatDesa(w http.ResponseWriter, r *http.Request) {
perangkats, err := ph.PerangkatService.PerangkatDesaRepository.GetAllPerangkatDesa()
if err != nil {
	http.Error(w, "Error retrieving perangkats", http.StatusInternalServerError)
	return
}

jsonResponse, err := json.Marshal(perangkats)
if err != nil {
	http.Error(w, "Internal server error", http.StatusInternalServerError)
	return
}

w.Header().Set("Content-Type", "application/json")
w.WriteHeader(http.StatusOK)
w.Write(jsonResponse)
}