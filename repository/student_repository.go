// Package repository contains methods for interacting with the database for student-related operations.
package repository

import (
	"github.com/HrithikSawant/go-crud-api/model"
	"gorm.io/gorm"
)

// StudentRepository handles CRUD operations for students.
type StudentRepository struct {
	DB *gorm.DB
}

// NewStudentRepository creates a new instance of StudentRepository.
func NewStudentRepository(db *gorm.DB) *StudentRepository {
	return &StudentRepository{DB: db}
}

// Create creates a new student in the database and returns the created student object.
func (r *StudentRepository) Create(student *model.Student) (*model.Student, error) {
	// Create the student in the DB
	if err := r.DB.Create(student).Error; err != nil {
		return nil, err
	}
	// Return the created student object
	return student, nil
}

// GetByID retrieves a student by its ID.
func (r *StudentRepository) GetByID(id uint) (*model.Student, error) {
	var student model.Student
	if err := r.DB.First(&student, id).Error; err != nil {
		return nil, err
	}
	return &student, nil
}

// Update updates an existing student in the database.
func (r *StudentRepository) Update(student *model.Student) (*model.Student, error) {
	if err := r.DB.Save(student).Error; err != nil {
		return nil, err
	}
	return student, nil
}

// Delete deletes a student from the database.
func (r *StudentRepository) Delete(id uint) error {
	if err := r.DB.Delete(&model.Student{}, id).Error; err != nil {
		return err
	}
	return nil
}

// GetAll retrieves all students from the database.
func (r *StudentRepository) GetAll() ([]model.Student, error) {
	var students []model.Student
	if err := r.DB.Find(&students).Error; err != nil {
		return nil, err
	}
	return students, nil
}
