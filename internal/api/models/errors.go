package models

import "fmt"

// NotFoundError is a domain-layer error related to non-existent resources
type NotFoundError string

func (e NotFoundError) Error() string {
	if e == "" {
		return "Resource not found"
	}
	return fmt.Sprintf("Resource not found: %s", e)
}

