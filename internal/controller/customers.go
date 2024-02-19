package controller

import (
	"fmt"
	"net/http"
	"rinha-backend-2024q1-rafael-pissolatto-nunes/internal/database"
	"rinha-backend-2024q1-rafael-pissolatto-nunes/internal/repository"
	"rinha-backend-2024q1-rafael-pissolatto-nunes/internal/util"
	"strconv"

	"github.com/gorilla/mux"
)

// GetCustomers is a function that returns a list of customers
func GetCustomers(w http.ResponseWriter, r *http.Request) {
	db, err := database.Connect()
	if err != nil {
		util.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repository.NewCustomersRepository(db)
	customers, err := repository.GetCustomers()
	if err != nil {
		util.Error(w, http.StatusInternalServerError, err)
		return
	}

	if len(customers) == 0 {
		util.Error(w, http.StatusNotFound, fmt.Errorf("No customers found"))
		return
	}

	util.JSON(w, http.StatusOK, customers)
}

// getIDFromRequest is a function that returns the ID from the request, ex /clientes/{id}/extrato
func getIDFromRequest(r *http.Request) (int, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return 0, fmt.Errorf("ID not found in request")
	}

	customerID, err := strconv.Atoi(id)
	if err != nil {
		return 0, fmt.Errorf("Invalid ID")
	}

	return customerID, nil
}

// GetStatementByCustomerId is a function that returns a list of transactions for a customer
func GetStatementByCustomerId(w http.ResponseWriter, r *http.Request) {
	db, err := database.Connect()
	if err != nil {
		util.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repository.NewCustomersRepository(db)
	customerID, err := getIDFromRequest(r)
	if err != nil {
		util.Error(w, http.StatusBadRequest, err)
		return
	}

	statement, err := repository.GetStatementByCustomerId(customerID)
	if err != nil {
		util.Error(w, http.StatusInternalServerError, err)
		return
	}

	util.JSON(w, http.StatusOK, statement)
}

// PostTransaction is a function that creates a new transaction for a customer
// func PostTransaction(w http.ResponseWriter, r *http.Request) {
// 	db, err := database.Connect()
// 	if err != nil {
// 		util.Error(w, http.StatusInternalServerError, err)
// 		return
// 	}
// 	defer db.Close()

// 	repository := repository.NewCustomersRepository(db)
// 	customerID, err := getIDFromRequest(r)
// 	if err != nil {
// 		util.Error(w, http.StatusBadRequest, err)
// 		return
// 	}

// 	transaction, err := util.DecodeTransaction(r)
// 	if err != nil {
// 		util.Error(w, http.StatusBadRequest, err)
// 		return
// 	}

// 	err = repository.CreateTransaction(customerID, transaction)
// 	if err != nil {
// 		util.Error(w, http.StatusInternalServerError, err)
// 		return
// 	}

// 	util.JSON(w, http.StatusCreated, nil)
// }
