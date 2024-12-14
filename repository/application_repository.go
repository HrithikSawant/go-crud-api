// Package repository provides functions for interacting with the database for application-related entities.
package repository

import (
	"gorm.io/gorm"
)

// ApplicationRepository will manage all repositories, including the student repository.
type ApplicationRepository struct {
	StudentRepository *StudentRepository
	HealthRepository  *HealthRepository
}

// NewApplicationRepository initializes the ApplicationRepository with the necessary dependencies.
func NewApplicationRepository(db *gorm.DB) (*ApplicationRepository, error) {
	// Initialize StudentRepository
	studentRepo := NewStudentRepository(db)
	// Initialize HealthRepository
	healthRepo := NewHealthRepository(db)

	// Return the ApplicationRepository that includes student repository.
	return &ApplicationRepository{
		StudentRepository: studentRepo,
		HealthRepository:  healthRepo,
	}, nil
}
