package repository

import (
	"database/sql"
	"net/http"

	"github.com/faruan/desafio-go-imersao14/internal/freight/entity"
)

type RouteRepositoryMysql struct {
	db *sql.DB
}

func NewRouteRepositoryMysql(db *sql.DB) *RouteRepositoryMysql {
	return &RouteRepositoryMysql{
		db: db,
	}
}

func (r *RouteRepositoryMysql) ListRoutes(w http.ResponseWriter, req *http.Request) {
	rows, err := r.db.Query("SELECT * FROM routes")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
}

func (r *RouteRepositoryMysql) Create(route *entity.Route) error {
	sql := "INSERT INTO routes (name, source_lat, source_lng, destination_lat, destination_lng) VALUES (?, ?, ?, ?, ?)"
	_, err := r.db.Exec(sql, route.Name, route.Source.Lat, route.Source.Lng, route.Destination.Lat, route.Destination.Lng)
	if err != nil {
		return err
	}
	return nil
}
