package storage

import (
	"database/sql"
	"errors"

	"github.com/gdlroutes/api/internal/api/models"
	"github.com/gdlroutes/api/internal/api/usecases/geodata"
)

type postgresStorage struct {
	db *sql.DB
}

var _ geodata.Storage = &postgresStorage{}

// NewPostgres returns a new Postgres storage
func NewPostgres(db *sql.DB) (geodata.Storage, error) {
	if db == nil {
		return nil, errors.New("nil db")
	}

	return &postgresStorage{db}, nil
}

func (s *postgresStorage) DoesCategoryExist(categoryID int) (bool, error) {
	const query = `
	SELECT id
	FROM categories
	WHERE id = $1;
	`
	var id int
	switch err := s.db.QueryRow(query, categoryID).Scan(&id); err {
	case nil:
		return true, nil
	case sql.ErrNoRows:
		return false, nil
	default:
		return false, err
	}
}

func (s *postgresStorage) getHotspotsByCategory(categoryID int) ([]*models.Hotspot, error) {
	const query = `
		SELECT id, name, description, type, latitude, longitude
		FROM hotspots
		WHERE category_id = $1;
		`
	rows, err := s.db.Query(query, categoryID)
	if err != nil {
		return nil, err
	}

	hotspots := make([]*models.Hotspot, 0)
	for rows.Next() {
		hotspot := &models.Hotspot{
			Coords: [2]float64{},
		}
		var description sql.NullString
		if err := rows.Scan(
			&hotspot.ID,
			&hotspot.Name,
			&description,
			&hotspot.Type,
			&hotspot.Coords[0],
			&hotspot.Coords[1],
		); err != nil {
			return nil, err
		}
		hotspot.Description = description.String
		hotspots = append(hotspots, hotspot)
	}

	return hotspots, nil
}

func (s *postgresStorage) GetCategories() ([]*models.Category, error) {
	const query = `
	SELECT id, name
	FROM categories;
	`
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}

	categories := make([]*models.Category, 0)
	for rows.Next() {
		category := &models.Category{
			Active:   true,
			Hotspots: make([]*models.Hotspot, 0),
		}
		if err := rows.Scan(&category.ID, &category.Name); err != nil {
			return nil, err
		}

		category.Hotspots, err = s.getHotspotsByCategory(category.ID)
		if err != nil {
			return nil, err
		}

		categories = append(categories, category)
	}

	return categories, nil
}

func (s *postgresStorage) GetCategoryByID(categoryID int) (*models.Category, error) {
	category := &models.Category{Active: true}

	const query = `
	SELECT id, name
	FROM categories
	WHERE id = $1;
	`
	if err := s.db.QueryRow(query, categoryID).Scan(&category.ID, &category.Name); err != nil {
		return nil, err
	}

	var err error
	category.Hotspots, err = s.getHotspotsByCategory(category.ID)
	if err != nil {
		return nil, err
	}

	return category, nil
}

func (s *postgresStorage) GetRouteCategories() ([]*models.RouteCategory, error) {
	const query = `
	SELECT id, name
	FROM route_categories;
	`
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}

	categories := make([]*models.RouteCategory, 0)
	for rows.Next() {
		category := &models.RouteCategory{}
		if err := rows.Scan(&category.ID, &category.Name); err != nil {
			return nil, err
		}

		categories = append(categories, category)
	}

	return categories, nil
}

func createRoutePoints(tx *sql.Tx, routeID int, points [][2]float64) error {
	const query = `
	INSERT INTO route_points(route_id, latitude, longitude)
	VALUES($1, $2, $3);
	`
	for _, point := range points {
		if _, err := tx.Exec(query, routeID, point[0], point[1]); err != nil {
			return err
		}
	}

	return nil
}

func (s *postgresStorage) CreateRoute(route *models.Route) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	const query = `
	INSERT INTO routes(name, description, category_id)
	VALUES ($1, $2, $3)
	RETURNING id;
	`
	var id int
	if err := tx.QueryRow(query, route.Name, route.Description, route.CategoryID).Scan(&id); err != nil {
		tx.Rollback()
		return err
	}

	if err := createRoutePoints(tx, id, route.Points); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (s *postgresStorage) getRoutePoints(routeID int) ([][2]float64, error) {
	const query = `
	SELECT latitude, longitude
	FROM route_points
	JOIN routes
	ON route_points.route_id = routes.id
	WHERE routes.id = $1
	ORDER BY route_points.id ASC;
	`
	rows, err := s.db.Query(query, routeID)
	if err != nil {
		return nil, err
	}

	points := make([][2]float64, 0)
	for rows.Next() {
		point := [2]float64{}
		if err := rows.Scan(&point[0], &point[1]); err != nil {
			return nil, err
		}
		points = append(points, point)
	}

	return points, nil
}

func (s *postgresStorage) GetAllRoutes() ([]*models.Route, error) {
	const query = `
	SELECT routes.id, routes.name, routes.description, categories.id
	FROM routes
	JOIN categories
	ON routes.category_id = categories.id
	`
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}

	routes := make([]*models.Route, 0)
	for rows.Next() {
		var description sql.NullString
		route := &models.Route{}
		if err := rows.Scan(&route.ID, &route.Name, &description, &route.CategoryID); err != nil {
			return nil, err
		}
		route.Description = description.String

		route.Points, err = s.getRoutePoints(route.ID)
		if err != nil {
			return nil, err
		}

		routes = append(routes, route)
	}

	return routes, nil
}

func (s *postgresStorage) GetRoutesByCategory(categoryID int) ([]*models.Route, error) {
	const query = `
	SELECT routes.id, routes.name, routes.description, categories.id
	FROM routes
	JOIN categories
	ON routes.category_id = categories.id
	WHERE categories.id = $1;
	`
	rows, err := s.db.Query(query, categoryID)
	if err != nil {
		return nil, err
	}

	routes := make([]*models.Route, 0)
	for rows.Next() {
		var description sql.NullString
		route := &models.Route{}
		if err := rows.Scan(&route.ID, &route.Name, &description, &route.CategoryID); err != nil {
			return nil, err
		}
		route.Description = description.String

		route.Points, err = s.getRoutePoints(route.ID)
		if err != nil {
			return nil, err
		}

		routes = append(routes, route)
	}

	return routes, nil
}
