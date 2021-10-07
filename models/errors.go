package models

type ErrorResponse struct {
	Message string `json:"message,string"`
}

type ValidationErrorResponse struct {
	Message string            `json:"message"`
	Errors  map[string]string `json:"errors,string"`
}
