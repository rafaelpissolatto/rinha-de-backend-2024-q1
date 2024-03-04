package routes

import (
	"net/http"
	"rinha-backend-2024q1-rafael-pissolatto-nunes/internal/controller"
)

var routeCustomers = []Route{
	{
		URI:                     "/clientes",
		Method:                  http.MethodGet,
		Function:                controller.GetCustomers,
		RequireAuthentification: false,
	},
	{
		URI:                     "/clientes/{id}/extrato",
		Method:                  http.MethodGet,
		Function:                controller.GetCompleteStatementByCustomerId,
		RequireAuthentification: false,
	},
	{
		URI:                     "/clientes/{id}/transacoes",
		Method:                  http.MethodPost,
		Function:                controller.PostTransactionByCustomerId,
		RequireAuthentification: false,
	},
}
