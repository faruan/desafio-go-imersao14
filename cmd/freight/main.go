package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/faruan/desafio-go-imersao14/internal/freight/entity"
	"github.com/faruan/desafio-go-imersao14/internal/freight/infra/repository"
	"github.com/go-chi/chi"
)

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("mysql", "root:root@tcp(host.docker.internal:3306)/routes?parseTime=true")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	routeRepo := repository.NewRouteRepositoryMysql(db)

	router := chi.NewRouter()

	router.Get("/api/routes", routeRepo.ListRoutes)
	router.Post("/api/routes", func(w http.ResponseWriter, r *http.Request) {
		var route entity.Route
		err := json.NewDecoder(r.Body).Decode(&route)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = routeRepo.Create(&route)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	})

	fmt.Println("Server is running on :8080")
	http.ListenAndServe(":8080", router)
}
