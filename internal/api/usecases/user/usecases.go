package user

import "github.com/gdlroutes/api/internal/api/models"

// UseCases is a domain-layer set of functions related to users
type UseCases interface {
	CreateUserAndToken(user *models.User) (*models.Token, error)
	CreateToken(user *models.User) (*models.Token, error)
}
