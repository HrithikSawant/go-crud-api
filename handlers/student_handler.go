// Package handlers provides HTTP handlers for managing student-related operations.
package handlers

import (
	"net/http"
	"strconv"

	"github.com/HrithikSawant/go-crud-api/model"
	service "github.com/HrithikSawant/go-crud-api/services"
	"github.com/gin-gonic/gin"
)

// StudentHandler handles HTTP requests related to student operations.
type StudentHandler struct {
	studentService service.StudentService
}

// NewStudentHandler creates a new instance of StudentHandler.
func NewStudentHandler(studentService *service.StudentService) *StudentHandler {
	return &StudentHandler{studentService: *studentService}
}

// RegisterRoutes registers all routes for the student resource.
func (h *StudentHandler) RegisterRoutes(r *gin.Engine) {
	studentRoutes := r.Group("/students")
	{
		studentRoutes.POST("/", h.CreateStudent)
		studentRoutes.GET("/:id", h.GetStudentByID)
		studentRoutes.PUT("/:id", h.UpdateStudent)
		studentRoutes.DELETE("/:id", h.DeleteStudent)
		studentRoutes.GET("/", h.GetAllStudents)
	}
}

// CreateStudent handles the creation of a new student.
func (h *StudentHandler) CreateStudent(c *gin.Context) {
	var student model.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdStudent, err := h.studentService.CreateStudent(&student)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdStudent)
}

// GetStudentByID retrieves a student by its ID.
func (h *StudentHandler) GetStudentByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32) // Convert the string ID to uint
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	student, err := h.studentService.GetStudentByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}
	c.JSON(http.StatusOK, student)
}

// UpdateStudent updates an existing student by ID.
func (h *StudentHandler) UpdateStudent(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32) // Convert the string ID to uint
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var student model.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	student.ID = uint(id) // Assign the parsed ID
	updatedStudent, err := h.studentService.UpdateStudent(&student)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedStudent)
}

// DeleteStudent deletes a student by ID.
func (h *StudentHandler) DeleteStudent(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32) // Convert the string ID to uint
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	err = h.studentService.DeleteStudent(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

// GetAllStudents retrieves all students.
func (h *StudentHandler) GetAllStudents(c *gin.Context) {
	students, err := h.studentService.GetAllStudents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, students)
}
