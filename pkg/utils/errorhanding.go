package util

import (
	"log"
)

// CustomError is a custom error structure that wraps an error message and additional data.
type CustomError struct {
	Message string
	Data    map[string]interface{}
}

// Error returns the error message of the CustomError.
func (e *CustomError) Error() string {
	return e.Message
}

// NewError creates a new CustomError with the given message and optional data.
func NewError(message string, data map[string]interface{}) error {
	return &CustomError{
		Message: message,
		Data:    data,
	}
}

// LogError logs the error message and any additional data.
func LogError(err error) {
	if err != nil {
		log.Printf("Error: %v\n", err)
		if customErr, ok := err.(*CustomError); ok {
			for key, value := range customErr.Data {
				log.Printf("%s: %v\n", key, value)
			}
		}
	}
}
