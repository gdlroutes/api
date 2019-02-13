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
				Name:   "Libraries",
				Active: true,
				Hotspots: []*models.Hotspot{
					&models.Hotspot{
						ID:          13,
						Name:        "Biblioteca Central Estatal Profesor Ramón García Ruiz",
						Description: "Antiguamente Capilla de la Preciosa Sangre, en 1978 se convirtió en bodega de libros del Gobierno del Estado. En 1981 el inmueble fue restaurado para convertirse en biblioteca y actualmente cuenta con salas de consulta, general e infantil, asícomo el fondo especial de escritores jaliscienses.",
						Type:        "marker",
						Coords:      [2]float64{20.6868385, -103.35037},
					},
					&models.Hotspot{
						ID:          14,
						Name:        "Biblioteca Iberoamericana Octavio Paz",
						Description: "Fundado en 1561 por los jesuitas, albergó a la Universidad Neogalaica de 1793 a 1827. Fungió como Palacio Legislativo, luego como universidad y oficina de Telégrafos; tras 6 años de abandono, en 1991 se inauguró como sede de la biblioteca, honrando el nombre del Nobel de Literatura Octavio Paz.",
						Type:        "marker",
						Coords:      [2]float64{20.6756959, -103.3483031},
					},
					&models.Hotspot{
						ID:          15,
						Name:        "Biblioteca del Ejército y Fuerza Aérea Mexicana",
						Description: "Conocida en 1821 como 'La Casa de la Campana', pues aquí se fundió la campana mayor de Catedral, fue fábrica de pólvora durante la Reforma; se clausuró al poco tiempo por una explosión que casi la destruye. Desde 1999, alberga más de 40,000 libros y documentos, así como una sala para invidentes.",
						Type:        "marker",
						Coords:      [2]float64{20.6862386, -103.3511596},
					},
					&models.Hotspot{
						ID:          16,
						Name:        "Biblioteca Pública del Estado de Jalisco \"Juan José Arreola\"",
						Description: "La biblioteca se fundó el 24 de julio de 1861 y fue inaugurada el 18 de diciembre de 1874. En 1925 pasó a manos de la Universidad de Guadalajara y al día de hoy ha contado con cuatro sedes. En 2001 fue bautizada con el nombre de 'Juan José Arreola', quien fuera su director durante diez años.",
						Type:        "marker",
						Coords:      [2]float64{20.738445, -103.3813623},
					},
				},
			},
			1: &models.Category{
				ID:     1,
				Name:   "Bosques",
				Active: true,
				Hotspots: []*models.Hotspot{
					&models.Hotspot{
						ID:          17,
						Name:        "Área de Protección de Flora y Fauna La Primavera",
						Description: "Su nombre se debe a la Hacienda La Primavera, construida en 1923. Con 140,000 años de existencia, es la más grande reserva ecológica que le queda a la zona metropolitana. Se ubica en la confluencia de la Sierra Madre Occidental y el Eje Neovolcánico Transversal. En 1980 se declaró zona protegida.",
						Type:        "marker",
						Coords:      [2]float64{20.7011129, -103.5781066},
					},
					&models.Hotspot{
						ID:          18,
						Name:        "Bosque El Centinela",
						Description: "El bosque original que existió en este sitio hasta el siglo XVIII fue devastado para uso de combustible, dada su cercanía con Guadalajara. Las primeras plantaciones para restaurar El Centinela se llevaron a cabo en la década de 1970, usando 80% de especies exóticas para asegurar su supervivencia.",
						Type:        "marker",
						Coords:      [2]float64{20.760076, -103.3823272},
					},
					&models.Hotspot{
						ID:          19,
						Name:        "Bosque Los Colomos (Ingreso El Chaco)",
						Description: "De sus 92 hectáreas de extensión, en este ingreso al bosque se encuentra la zona de eucaliptos, un área de esparcimiento familiar; también se puede encontrar El Castillo, que funciona como centro cultural, y el monumento a Pepe Guízar, una estatua del afamado compositor de la música vernácula.",
						Type:        "marker",
						Coords:      [2]float64{20.7040375, -103.3899401},
					},
					&models.Hotspot{
						ID:          19,
						Name:        "Bosque Los Colomos (Ingreso Paseo Torreón)",
						Description: "Entre los 32,000 ejemplares de árboles, cerca de este ingreso se encuentra el Jardín Japonés, un santuario solemne donado por la ciudad de Kyoto; asimismo, el Lago de Aves es un espejo de agua artificial para varias especies de aves migratorias. Este bosque aún abastece de agua a las zonas aledañas.",
						Type:        "marker",
						Coords:      [2]float64{20.7060597, -103.3939649},
					},
				},
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
