package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"

	"github.com/Yoas-Hutapea/Microservice_09/pengumuman/models"
	"github.com/Yoas-Hutapea/Microservice_09/pengumuman/services"
)

type PengumumanHandler struct {
	PengumumanService *services.PengumumanService
}

func NewPengumumanHandler(pengumumanService *services.PengumumanService) *PengumumanHandler {
	// Initialize and configure the PendudukHandler instance
	return &PengumumanHandler{
		PengumumanService: pengumumanService,
		// Initialize other fields or dependencies
	}
}

func (ph *PengumumanHandler) AddPengumuman(w http.ResponseWriter, r *http.Request) {
	// Parse the request body
	var pengumuman models.Pengumuman
	err := json.NewDecoder(r.Body).Decode(&pengumuman)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Call the add pengumuman service
	err = ph.PengumumanService.AddPengumuman(&pengumuman)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return success response
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, "Pengumuman added successfully")
}

func (ph *PengumumanHandler) UpdatePengumuman(w http.ResponseWriter, r *http.Request) {
	// Parse the request body
	var pengumuman models.Pengumuman
	err := json.NewDecoder(r.Body).Decode(&pengumuman)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Call the update pengumuman service
	err = ph.PengumumanService.UpdatePengumuman(&pengumuman)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return success response
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Pengumuman updated successfully")
}

func (ph *PengumumanHandler) DeletePengumuman(w http.ResponseWriter, r *http.Request) {
	// Parse the request parameters
	// Assuming the pengumuman ID is passed as a query parameter named "id"
	pengumumanID := r.URL.Query().Get("id")
	if pengumumanID == "" {
		http.Error(w, "Pengumuman ID is required", http.StatusBadRequest)
		return
	}

	// Convert pengumumanID to int
	id, err := strconv.Atoi(pengumumanID)
	if err != nil {
		http.Error(w, "Invalid Pengumuman ID", http.StatusBadRequest)
		return
	}

	// Call the delete pengumuman service
	err = ph.PengumumanService.DeletePengumuman(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return success response
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Pengumuman deleted successfully")
}

func (ph *PengumumanHandler) GetPengumumanByID(w http.ResponseWriter, r *http.Request) {
	// Parse the request parameters
	// Assuming the pengumuman ID is passed as a path variable named "id"
	vars := mux.Vars(r)
	pengumumanID := vars["id"]
	if pengumumanID == "" {
		http.Error(w, "Pengumuman ID is required", http.StatusBadRequest)
		return
	}

	// Convert pengumumanID to int
	id, err := strconv.Atoi(pengumumanID)
	if err != nil {
		http.Error(w, "Invalid Pengumuman ID", http.StatusBadRequest)
		return
	}

	// Call the get pengumuman by ID service
	pengumuman, err := ph.PengumumanService.GetPengumumanByID(id)
	if err != nil {
		http.Error(w, "Error retrieving pengumuman", http.StatusInternalServerError)
		return
	}

	// Marshal pengumuman to JSON
	jsonResponse, err := json.Marshal(pengumuman)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Set response headers and write JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}


func (ph *PengumumanHandler) GetAllPengumuman(w http.ResponseWriter, r *http.Request) {
	pengumumans, err := ph.PengumumanService.PengumumanRepository.GetAllPengumuman()
	if err != nil {
		http.Error(w, "Error retrieving pengumumans", http.StatusInternalServerError)
		return
	}

	jsonResponse, err := json.Marshal(pengumumans)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}