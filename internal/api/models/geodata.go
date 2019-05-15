package models

// Hotspot is a hotspot object
type Hotspot struct {
	ID          int        `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Type        string     `json:"type"`
	Coords      [2]float64 `json:"coords"`
}

// Category is a category object
type Category struct {
	ID       int        `json:"id"`
	Name     string     `json:"name"`
	Active   bool       `json:"active"`
	Hotspots []*Hotspot `json:"features"`
}

// Route is a route object
type Route struct {
	ID          int          `json:"id"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	CategoryID  int          `json:"category_id"`
	Points      [][2]float64 `json:"points"`
}

// RouteCategory is a category for routes
type RouteCategory struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
