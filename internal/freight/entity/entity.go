package entity

type RouteRepository interface {
	Create(route *Route) error
	FindAll() (*Route, error)
}

type Route struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Source      Place  `json:"source"`
	Destination Place  `json:"destination"`
}

type Place struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}
