package routes

import (
	"net/http"

	"rinha-backend-2024q1-rafael-pissolatto-nunes/internal/controller"
)

var routeMetrics = Route{
	URI:                     "/metrics",
	Method:                  http.MethodGet,
	Function:                controller.Metrics,
	RequireAuthentification: false,
}
