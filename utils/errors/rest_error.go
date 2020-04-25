package errors

import "net/http"

type RestError struct {
	Message string `json:"message"`
	Status  int    `json:"code"`
	Error   string `json:"error"`
}

func BadRequestError(message string) *RestError {
	return &RestError{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "bad request",
	}
}

func UserAlreadyPresent(message string) *RestError {
	return &RestError{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "User already present",
	}
}
func UserNotFoundError(message string) *RestError {
	return &RestError{
		Message: message,
		Status:  http.StatusBadRequest,
	}
}
