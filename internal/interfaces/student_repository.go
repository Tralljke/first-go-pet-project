package interfaces

import "first-go-pet-project/internal/entities"

type StudentRepository interface {
	GetAll() ([]entities.Student, error)
}
