package geodata

import "github.com/gdlroutes/api/internal/api/models"

// Storage is a generic interface for a geodata storage
type Storage interface {
	DoesCategoryExist(categoryID int) (bool, error)
	GetCategories() ([]*models.Category, error)
	GetCategoryByID(categoryID int) (*models.Category, error)
	GetRouteCategories() ([]*models.RouteCategory, error)
	CreateRoute(route *models.Route) error
	GetAllRoutes() ([]*models.Route, error)
	GetRoutesByCategory(categoryID int) ([]*models.Route, error)
}
