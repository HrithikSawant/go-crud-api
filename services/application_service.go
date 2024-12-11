package service

import "github.com/HrithikSawant/go-crud-api/repository"

// ApplicationService aggregates various services (e.g., StudentService, HealthService)
// to provide high-level application operations.
type ApplicationService struct {
	StudentService *StudentService
	HealthService  *HealthService
	// You can add more services like CoursesService, etc.
}

// NewApplicationService initializes the ApplicationService with necessary repositories and services.
func NewApplicationService(applicationRepository *repository.ApplicationRepository) *ApplicationService {
	return &ApplicationService{
		StudentService: newStudentService(applicationRepository.StudentRepository),
		HealthService:  newHealthService(applicationRepository.HealthRepository),
	}
}
