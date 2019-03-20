package user

import "github.com/gdlroutes/api/internal/api/models"

// Storage is a generic interface for a user storage
type Storage interface {
	GetUserByEmail(email string) (*models.User, error)
	CreateUser(user *models.User) (string, error)
	IsPasswordCorrect(email, password string) (bool, error)
}
