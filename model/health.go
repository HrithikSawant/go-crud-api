package model

// HealthStatus represents the health status of the application, including its message and state.
type HealthStatus struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
