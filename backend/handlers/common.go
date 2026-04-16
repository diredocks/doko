package handlers

type Response struct {
	Success bool   `json:"success" example:"true"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

type ErrorResponse struct {
	Success bool   `json:"success" example:"false"`
	Message string `json:"message"`
}
