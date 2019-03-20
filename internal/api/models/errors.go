package models

import "fmt"

// NotFoundError is a domain-layer error related to non-existent resources
type NotFoundError string

func (e NotFoundError) Error() string {
	if e == "" {
		return "Resource not found"
	}
	return fmt.Sprintf("Resource not found: %s", string(e))
}

// ConflictError is a domain-layer error related to already existing unique resources
type ConflictError string

func (e ConflictError) Error() string {
	if e == "" {
		return "Conflict"
	}
	return fmt.Sprintf("Conflict: %s", string(e))
}

// InvalidCredentialsError is a domain-layer error related to bad credentials
type InvalidCredentialsError string

func (e InvalidCredentialsError) Error() string {
	if e == "" {
		return "Invalid credentials"
	}
	return fmt.Sprintf("Invalid credentials: %s", string(e))
}
