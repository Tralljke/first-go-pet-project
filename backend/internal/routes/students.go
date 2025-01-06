package routes

import (
	"first-go-pet-project/backend/internal/handlers"
	"first-go-pet-project/backend/internal/usecases"
	"net/http"
)

func RegisterRoutes(mux *http.ServeMux, studentService *usecases.StudentService) {
	studentHandler := handlers.NewStudentHandler(studentService)

	mux.HandleFunc("/students", studentHandler.GetStudents)
}
