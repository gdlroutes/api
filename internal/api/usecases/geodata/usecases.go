package geodata

import "github.com/gdlroutes/api/internal/api/models"

// UseCases is a domain-layer set of functions related to hotspots
type UseCases interface {
	GetCategories() ([]*models.Category, error)
	GetCategoryByID(categoryID int) ([]*models.Category, error)
}
