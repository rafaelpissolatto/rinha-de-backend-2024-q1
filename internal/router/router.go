package router

import (
	"github.com/gorilla/mux"

	"rinha-backend-2024q1-rafael-pissolatto-nunes/internal/router/routes"
)

func Generate() *mux.Router {
	r := mux.NewRouter()
	return routes.Configure(r)
}
