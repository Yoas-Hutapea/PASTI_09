package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Yoas-Hutapea/Microservice_09/api/models"
	"github.com/Yoas-Hutapea/Microservice_09/api/services"
)

type PengumumanHandler struct {
	PengumumanService *services.PengumumanService
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

	// Call the delete pengumuman service
	err := ph.PengumumanService.DeletePengumuman(pengumumanID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return success response
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Pengumuman deleted successfully")
}