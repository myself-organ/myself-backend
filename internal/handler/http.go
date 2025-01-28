package handler

import (
	"encoding/json"
	"myself-backend/internal/domain"
	"myself-backend/internal/service"
	"net/http"
	"strconv"
)

type CVHandler struct {
	service *service.CVService
}

func NewCVHandler(service *service.CVService) *CVHandler {
	return &CVHandler{service: service}
}

func (h *CVHandler) CreateCV(w http.ResponseWriter, r *http.Request) {
	var cv domain.CV
	if err := json.NewDecoder(r.Body).Decode(&cv); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.service.CreateCV(cv); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *CVHandler) GetCVByID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "missing id parameter", http.StatusBadRequest)
		return
	}

	cvID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "invalid id parameter", http.StatusBadRequest)
		return
	}

	cv, err := h.service.GetCVByID(cvID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(cv)
}

func (h *CVHandler) GetAllCVs(w http.ResponseWriter, r *http.Request) {
	cvs, err := h.service.GetAllCVs()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(cvs)
}

// ...other HTTP handlers...
