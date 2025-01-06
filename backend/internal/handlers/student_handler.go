package handlers

import (
	"encoding/json"
	"first-go-pet-project/backend/internal/usecases"
	"net/http"
)

type StudentHandler struct {
	service *usecases.StudentService
}

func NewStudentHandler(service *usecases.StudentService) *StudentHandler {
	return &StudentHandler{service: service}
}

func (h *StudentHandler) GetStudents(w http.ResponseWriter, r *http.Request) {
	students, err := h.service.GetAllStudents()
	if err != nil {
		http.Error(w, "Failed to fetch students", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(students); err != nil {
		http.Error(w, "Failed to encode students", http.StatusInternalServerError)
	}
}
