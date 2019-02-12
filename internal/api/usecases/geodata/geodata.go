package geodata

import "github.com/gdlroutes/api/internal/api/models"

type useCases struct{}

var _ UseCases = &useCases{}

// New creates a default set of GeoData use cases
func New() (UseCases, error) {
	return &useCases{}, nil
}

// TODO: implement
func (u *useCases) GetCategories() ([]*models.Category, error) {
	return nil, nil
}

// TODO: implement
func (u *useCases) GetCategoryByID(categoryID int) ([]*models.Category, error) {
	return nil, nil
}
