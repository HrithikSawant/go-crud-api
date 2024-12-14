// Package service contains the business logic and operations for managing student-related tasks.
package service

import (
	"github.com/HrithikSawant/go-crud-api/model"
	"github.com/HrithikSawant/go-crud-api/repository"
)

// StudentService provides business logic related to student operations.
type StudentService struct {
	studentRepo *repository.StudentRepository
}

// NewStudentService initializes a new StudentService with the provided StudentRepository.
func newStudentService(studentRepo *repository.StudentRepository) *StudentService {
	return &StudentService{
		studentRepo: studentRepo,
	}
}

// CreateStudent handles the creation of a new student.
func (s *StudentService) CreateStudent(student *model.Student) (*model.Student, error) {
	return s.studentRepo.Create(student)
}

// GetStudentByID retrieves a student by ID.
func (s *StudentService) GetStudentByID(id uint) (*model.Student, error) {
	return s.studentRepo.GetByID(id)
}

// GetAllStudents retrieves all students.
func (s *StudentService) GetAllStudents() ([]model.Student, error) {
	return s.studentRepo.GetAll()
}

// UpdateStudent updates an existing student's data.
func (s *StudentService) UpdateStudent(student *model.Student) (*model.Student, error) {
	return s.studentRepo.Update(student)
}

// DeleteStudent deletes a student by ID.
func (s *StudentService) DeleteStudent(id uint) error {
	return s.studentRepo.Delete(id)
}
