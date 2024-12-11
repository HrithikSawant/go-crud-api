package repository

import (
	"fmt"

	"github.com/HrithikSawant/go-crud-api/model"
	"gorm.io/gorm"
)

// HealthRepository is responsible for checking the health of the system
type HealthRepository struct {
	DB *gorm.DB // Database connection (you can extend this for other dependencies)
}

// NewHealthRepository initializes the HealthRepository
func NewHealthRepository(db *gorm.DB) *HealthRepository {
	return &HealthRepository{DB: db}
}

// CheckDatabaseHealth checks if the database is reachable and returns a Health model with an error.
func (r *HealthRepository) CheckDatabaseHealth() (*model.HealthStatus, error) {
	// Try executing a simple query to verify the database connection
	if err := r.DB.Exec("SELECT 1").Error; err != nil {
		// If the query fails, return an unhealthy health model and the error
		return &model.HealthStatus{
			Status:  "unhealthy",
			Message: fmt.Sprintf("database connection failed: %v", err),
		}, err
	}
	// If the query succeeds, return a healthy health model and no error
	return &model.HealthStatus{
		Status:  "healthy",
		Message: "database is reachable",
	}, nil
}
