package storage

import (
	"fmt"

	"github.com/gdlroutes/api/internal/api/models"
	"github.com/gdlroutes/api/internal/api/usecases/geodata"
)

type fakeStorage struct {
	categories map[int]*models.Category
}

var _ geodata.Storage = &fakeStorage{}

// NewFake returns a new fake storage
func NewFake() (geodata.Storage, error) {
	return &fakeStorage{
		categories: map[int]*models.Category{
			0: &models.Category{
				ID:     0,
				Name:   "Restaurants",
				Active: true,
				Hotspots: []*models.Hotspot{
					&models.Hotspot{
						ID:     0,
						Name:   "Bogart's Smokehouse",
						Type:   "marker",
						Coords: [2]float64{38.6109607, -90.2050322},
					},
					&models.Hotspot{
						ID:     1,
						Name:   "Pappy's Smokehouse",
						Type:   "marker",
						Coords: [2]float64{38.6350008, -90.2261532},
					},
				},
			},
			1: &models.Category{
				ID:       0,
				Name:     "Museums",
				Active:   true,
				Hotspots: []*models.Hotspot{},
			},
		},
	}, nil
}

func (s *fakeStorage) DoesCategoryExist(categoryID int) (bool, error) {
	_, ok := s.categories[categoryID]
	return ok, nil
}

func (s *fakeStorage) GetCategories() ([]*models.Category, error) {
	categories := make([]*models.Category, 0)
	for _, category := range s.categories {
		categories = append(categories, category)
	}

	return categories, nil
}

func (s *fakeStorage) GetCategoryByID(categoryID int) (*models.Category, error) {
	category, ok := s.categories[categoryID]
	if !ok {
		return nil, fmt.Errorf("Category with ID %d not found", categoryID)
	}

	return category, nil
}
