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

func (u *useCases) GetCategories() ([]*models.Category, error) {
	return u.storage.GetCategories()
}

func (u *useCases) GetCategoryByID(categoryID int) (*models.Category, error) {
	if categoryExists, err := u.storage.DoesCategoryExist(categoryID); err != nil {
		return nil, err
	} else if !categoryExists {
		return nil, models.NotFoundError("")
	}

	return u.storage.GetCategoryByID(categoryID)
}
