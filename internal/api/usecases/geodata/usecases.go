package geodata

import (
	"github.com/gdlroutes/api/internal/api/models"
)

// UseCases is a domain-layer set of functions related to hotspots
type UseCases interface {
	GetCategories() ([]*models.Category, error)
	GetCategoryByID(categoryID int) (*models.Category, error)
	GetRouteCategories() ([]*models.RouteCategory, error)
	CreateRoute(route *models.Route) error
	GetRoutes() ([]*models.Route, error)
	GetRoutesByCategory(categoryID int) ([]*models.Route, error)
}
