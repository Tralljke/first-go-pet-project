package interfaces

import "first-go-pet-project/backend/internal/entities"

type StudentRepository interface {
	GetAll() ([]entities.Student, error)
}
