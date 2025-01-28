package handler

import (
	"encoding/json"
	"myself-backend/internal/domain"
	"myself-backend/internal/service"
	"net/http"
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

// ...other HTTP handlers...
