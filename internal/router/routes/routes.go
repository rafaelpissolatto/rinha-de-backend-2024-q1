package routes

import (
	"net/http"
	"rinha-backend-2024q1-rafael-pissolatto-nunes/internal/middleware"

	"github.com/gorilla/mux"
)

// Route is a struct that contains the information about a route
type Route struct {
	URI                     string
	Method                  string
	Function                func(w http.ResponseWriter, r *http.Request)
	RequireAuthentification bool
}

func Configure(r *mux.Router) *mux.Router {
	routes := []Route{}
	routes = append(routes, routeCustomers...)

	for _, route := range routes {
		r.HandleFunc(route.URI,
			middleware.Logger(route.Function)).Methods(route.Method)
	}

	return r
}
