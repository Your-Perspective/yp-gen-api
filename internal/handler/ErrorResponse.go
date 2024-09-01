package handler

// ErrorResponse represents the standard error response
type ErrorResponse struct {
	Error   string            `json:"error"`
	Message string            `json:"message"`
	Fields  map[string]string `json:"fields,omitempty"` // This field will hold the detailed validation errors
}

// SuccessResponse represents a successful response message
type SuccessResponse struct {
	Message string `json:"message"`
}
