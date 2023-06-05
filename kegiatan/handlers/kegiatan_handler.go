package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"

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
	// Assuming the kegiatan ID is passed as a path variable named "id"
	vars := mux.Vars(r)
	kegiatanID := vars["id"]
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

func (kh *KegiatanHandler) GetKegiatanByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	kegiatanID, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid kegiatan ID", http.StatusBadRequest)
		return
	}

	kegiatan, err := kh.KegiatanService.KegiatanRepository.GetKegiatanByID(kegiatanID)
	if err != nil {
		http.Error(w, "Error retrieving kegiatan", http.StatusInternalServerError)
		return
	}

	jsonResponse, err := json.Marshal(kegiatan)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func (kh *KegiatanHandler) GetAllKegiatan(w http.ResponseWriter, r *http.Request) {
	kegiatans, err := kh.KegiatanService.KegiatanRepository.GetAllKegiatan()
	if err != nil {
		http.Error(w, "Error retrieving kegiatans", http.StatusInternalServerError)
		return
	}

	jsonResponse, err := json.Marshal(kegiatans)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}