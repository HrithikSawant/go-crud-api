package service

import (
	"github.com/HrithikSawant/go-crud-api/model"
	"github.com/HrithikSawant/go-crud-api/repository"
)

// HealthService provides business logic related to health checks.
type HealthService struct {
	healthRepo *repository.HealthRepository
}

// NewHealthService initializes a new HealthService with the provided dependencies.
func newHealthService(healthRepo *repository.HealthRepository) *HealthService {
	return &HealthService{
		healthRepo: healthRepo,
	}
}

// CheckHealth checks the health of the application (e.g., DB connectivity).
func (s *HealthService) CheckHealth() (*model.HealthStatus, error) {
	// Use HealthRepository to check database health and return both health model and error
	return s.healthRepo.CheckDatabaseHealth()
}
