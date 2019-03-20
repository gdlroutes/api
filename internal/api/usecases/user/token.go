package user

import (
	"time"

	"github.com/gdlroutes/api/internal/api/models"
)

// TokenGenerator is a generic interface for a user token generator
type TokenGenerator interface {
	GenerateToken(userID string, start time.Time) (*models.Token, error)
}
