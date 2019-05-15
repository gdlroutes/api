package geodata

import (
	"errors"

	"github.com/gdlroutes/api/internal/api/models"
)

type useCases struct {
	storage Storage
}

var _ UseCases = &useCases{}

// New creates a default set of geodata use cases
func New(storage Storage) (UseCases, error) {
	if storage == nil {
		return nil, errors.New("nil storage")
	}

	return &useCases{storage: storage}, nil
}

// GetCategories retrieves all categories
func (u *useCases) GetCategories() ([]*models.Category, error) {
	return u.storage.GetCategories()
}

// GetCategoryByID retrieves the category with given ID
func (u *useCases) GetCategoryByID(categoryID int) (*models.Category, error) {
	if categoryExists, err := u.storage.DoesCategoryExist(categoryID); err != nil {
		return nil, err
	} else if !categoryExists {
		return nil, models.NotFoundError("")
	}

	return u.storage.GetCategoryByID(categoryID)
}

// GetRouteCategories retrieves all route categories
func (u *useCases) GetRouteCategories() ([]*models.RouteCategory, error) {
	return u.storage.GetRouteCategories()
}

// CreateRoute creates a new route
func (u *useCases) CreateRoute(route *models.Route) error {
	if categoryExists, err := u.storage.DoesCategoryExist(route.CategoryID); err != nil {
		return err
	} else if !categoryExists {
		return models.NotFoundError("")
	}

	return u.storage.CreateRoute(route)
}

// GetRoutes fetches all routes
func (u *useCases) GetRoutes() ([]*models.Route, error) {
	return u.storage.GetAllRoutes()
}

// GetRoutesByCategory fetches all routes that match a given category
func (u *useCases) GetRoutesByCategory(categoryID int) ([]*models.Route, error) {
	return u.storage.GetRoutesByCategory(categoryID)
}
