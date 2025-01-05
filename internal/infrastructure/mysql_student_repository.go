package infrastructure

import (
	"first-go-pet-project/internal/entities"
	"gorm.io/gorm"
)

type MySQLStudentRepository struct {
	db *gorm.DB
}

func NewMySQLStudentRepository(db *gorm.DB) *MySQLStudentRepository {
	return &MySQLStudentRepository{db: db}
}

func (r *MySQLStudentRepository) GetAll() ([]entities.Student, error) {
	var students []entities.Student
	if err := r.db.Find(&students).Error; err != nil {
		return nil, err
	}
	return students, nil
}
