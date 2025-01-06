package usecases

import (
	"first-go-pet-project/backend/internal/entities"
	"first-go-pet-project/backend/internal/interfaces"
)

type StudentService struct {
	repo interfaces.StudentRepository
}

func NewStudentService(repo interfaces.StudentRepository) *StudentService {
	return &StudentService{repo: repo}
}

func (s *StudentService) GetAllStudents() ([]entities.Student, error) {
	return s.repo.GetAll()
}
