package helper

import "strings"

// Response - Common API Response Structure
type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Error   interface{} `json:"errors,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

// EmptyObj - If you don't wish to send data as nil then send this Empty Object
type EmptyObj struct{}

// BuildResponse method generate dynamic success response
func BuildResponse(mesaage string, data interface{}) Response {
	return Response{
		Success: true,
		Message: mesaage,
		Data:    data,
	}
}

// BuildErrorResponse method generate dynamic error Response
func BuildErrorResponse(mesaage string, error string) Response {
	errors := strings.Split(error, "\n")
	return Response{
		Success: false,
		Message: mesaage,
		Error:   errors,
	}
}
