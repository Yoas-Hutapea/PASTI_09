package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Yoas-Hutapea/Microservice_09/api/models"
	"github.com/Yoas-Hutapea/Microservice_09/api/services"
)

type KegiatanHandler struct {
	KegiatanService *services.KegiatanService
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

	// Call the delete kegiatan service
	err := kh.KegiatanService.DeleteKegiatan(kegiatanID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return success response
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Kegiatan deleted successfully")
}