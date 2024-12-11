package model

import (
	"errors"
	"regexp"
	"time"

	"gorm.io/gorm"
)

// Student represents the structure of a student entity.
type Student struct {
	ID        uint   `json:"id" gorm:"primaryKey"` // Primary Key for the Student entity
	FirstName string `json:"first_name"`           // First name of the student
	LastName  string `json:"last_name"`            // Last name of the student
	Email     string `json:"email" gorm:"unique"`  // Email address of the student (unique constraint)
	CreatedAt time.Time
	UpdatedAt time.Time
}

// BeforeCreate is a GORM hook that runs before creating a new record
func (student *Student) BeforeCreate(tx *gorm.DB) (err error) {
	// Validate if the email is not empty
	if student.Email == "" {
		return errors.New("email is required")
	}

	// Validate the email format using a regular expression
	emailRegex := `^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}$`
	if matched, _ := regexp.MatchString(emailRegex, student.Email); !matched {
		return errors.New("invalid email format")
	}

	// You can also check if the email already exists in the database (optional)
	var existingStudent Student
	if err := tx.Where("email = ?", student.Email).First(&existingStudent).Error; err == nil {
		return errors.New("email already exists")
	}

	return nil // Return nil to continue the creation process
}
