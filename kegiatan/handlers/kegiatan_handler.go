package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Yoas-Hutapea/Microservice_09/kegiatan/models"
	"github.com/Yoas-Hutapea/Microservice_09/kegiatan/services"
)

type KegiatanHandler struct {
	KegiatanService *services.KegiatanService
}

func NewKegiatanHandler(kegiatanService *services.KegiatanService) *KegiatanHandler {
	// Initialize and configure the PendudukHandler instance
	return &KegiatanHandler{
		KegiatanService: kegiatanService,
		// Initialize other fields or dependencies
	}
}

func (kh *KegiatanHandler) AddKegiatan(w http.ResponseWriter, r *http.Request) {
	// Parse the request body
	var kegiatan models.Kegiatan
	err := json.NewDecoder(r.Body).Decode(&kegiatan)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Call the add kegiatan service
	err = kh.KegiatanService.AddKegiatan(&kegiatan)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return success response
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, "Kegiatan added successfully")
}

func (kh *KegiatanHandler) UpdateKegiatan(w http.ResponseWriter, r *http.Request) {
	// Parse the request body
	var kegiatan models.Kegiatan
	err := json.NewDecoder(r.Body).Decode(&kegiatan)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Call the update kegiatan service
	err = kh.KegiatanService.UpdateKegiatan(&kegiatan)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return success response
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Kegiatan updated successfully")
}

func (kh *KegiatanHandler) DeleteKegiatan(w http.ResponseWriter, r *http.Request) {
	// Parse the request parameters
	// Assuming the kegiatan ID is passed as a query parameter named "id"
	kegiatanID := r.URL.Query().Get("id")
	if kegiatanID == "" {
		http.Error(w, "Kegiatan ID is required", http.StatusBadRequest)
		return
	}

	// Convert kegiatanID to int
	id, err := strconv.Atoi(kegiatanID)
	if err != nil {
		http.Error(w, "Invalid Kegiatan ID", http.StatusBadRequest)
		return
	}

	// Call the delete kegiatan service
	err = kh.KegiatanService.DeleteKegiatan(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return success response
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Kegiatan deleted successfully")
}
