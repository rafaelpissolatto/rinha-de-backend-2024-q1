package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"rinha-backend-2024q1-rafael-pissolatto-nunes/internal/database"
	"rinha-backend-2024q1-rafael-pissolatto-nunes/internal/model"
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

// GetCustomerByID is a function that returns a customer by its ID
func GetCustomerByID(w http.ResponseWriter, r *http.Request) {
	db, err := database.Connect()
	if err != nil {
		util.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repository.NewCustomersRepository(db)
	customerID, err := getCustomerIDFromRequest(r)
	if err != nil {
		util.Error(w, http.StatusBadRequest, err)
		return
	}

	customer, err := repository.GetCustomerByID(customerID)
	if err != nil {
		util.Error(w, http.StatusNotFound, err)
		return
	}

	util.JSON(w, http.StatusOK, customer)
}

// getCustomerIDFromRequest is a function that returns the ID from the request, ex /clientes/{id}/extrato
func getCustomerIDFromRequest(r *http.Request) (int, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return 0, fmt.Errorf("Customer ID not found in request")
	}

	customerID, err := strconv.Atoi(id)
	if err != nil {
		return 0, fmt.Errorf("Invalid Customer ID in request")
	}

	return customerID, nil
}

// decodeTransaction is a function that decodes the transaction from the request body
func decodeTransaction(r *http.Request) (model.Transaction, error) {
	var transaction model.Transaction
	err := json.NewDecoder(r.Body).Decode(&transaction)
	if err != nil {
		return model.Transaction{}, fmt.Errorf("Invalid request body")
	}

	return transaction, nil
}

// PostTransaction is a function that creates a new transaction for a customer
func PostTransactionByCustomerId(w http.ResponseWriter, r *http.Request) {
	db, err := database.Connect()
	if err != nil {
		util.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repository.NewCustomersRepository(db)
	customerID, err := getCustomerIDFromRequest(r)
	if err != nil {
		util.Error(w, http.StatusBadRequest, err)
		return
	}

	decodedRequest, err := decodeTransaction(r)
	if err != nil {
		util.Error(w, http.StatusBadRequest, err)
		return
	}

	// validate the transaction
	transaction := model.Transaction{
		CustomerID:   customerID,
		Amount:       decodedRequest.Amount,
		OperatorType: decodedRequest.OperatorType,
		Description:  decodedRequest.Description,
	}

	if err := transaction.Prepare(); err != nil {
		if err.Error() == "amount must be greater than or equal to zero" {
			util.Error(w, http.StatusUnprocessableEntity, err)
			return
		}
		if err.Error() == "description cannot be empty or greater than 10 characters" {
			util.Error(w, http.StatusUnprocessableEntity, err)
			return
		}
		if err.Error() == "customer id cannot be empty" {
			util.Error(w, http.StatusUnprocessableEntity, err)
			return
		}
		if err.Error() == "invalid operator type" {
			// 400 Bad Request
			util.Error(w, http.StatusBadRequest, err)
			return
		}

	// check if the customer exists
	if _, err := repository.GetCustomerByID(customerID); err != nil {
		util.Error(w, http.StatusNotFound, err)
		return
	}

	log.Println("[TRACE] TransactionDecoded: ", transaction)

	// create the transaction
	if err := repository.CreateTransaction(customerID, transaction); err != nil {
		//check if the error is from balance, return 422
		if err.Error() == "The customer's balance will be less than their available limit" {
			util.Error(w, http.StatusUnprocessableEntity, err)
			return
		}
	}

	// get the customer's statement
	statement, err := repository.GetSimpleStatementByCustomerId(customerID)
	if err != nil {
		util.Error(w, http.StatusInternalServerError, err)
		return
	}

	util.JSON(w, http.StatusOK, statement)
}

// GetSimpleStatementByCustomerId is a function that returns a simple statement for a customer
func GetSimpleStatementByCustomerId(w http.ResponseWriter, r *http.Request) {
	db, err := database.Connect()
	if err != nil {
		util.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repository.NewCustomersRepository(db)
	customerID, err := getCustomerIDFromRequest(r)
	if err != nil {
		util.Error(w, http.StatusBadRequest, err)
		return
	}

	// check if the customer exists
	if _, err := repository.GetCustomerByID(customerID); err != nil {
		util.Error(w, http.StatusNotFound, err)
		return
	}

	statement, err := repository.GetSimpleStatementByCustomerId(customerID)
	if err != nil {
		util.Error(w, http.StatusInternalServerError, err)
		return
	}

	util.JSON(w, http.StatusOK, statement)
}

// GetStatementByCustomerId is a function that returns a list of transactions for a customer
func GetCompleteStatementByCustomerId(w http.ResponseWriter, r *http.Request) {
	db, err := database.Connect()
	if err != nil {
		util.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repository.NewCustomersRepository(db)
	customerID, err := getCustomerIDFromRequest(r)
	if err != nil {
		util.Error(w, http.StatusBadRequest, err)
		return
	}

	// check if the customer exists
	if _, err := repository.GetCustomerByID(customerID); err != nil {
		util.Error(w, http.StatusNotFound, err)
		return
	}

	statement, err := repository.GetCompleteStatementByCustomerId(customerID)
	if err != nil {
		util.Error(w, http.StatusInternalServerError, err)
		return
	}

	util.JSON(w, http.StatusOK, statement)
}

// GetTransactions is a function that returns a list of transactions
func GetTransactions(w http.ResponseWriter, r *http.Request) {
	db, err := database.Connect()
	if err != nil {
		util.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repository.NewTransactionsRepository(db)
	transactions, err := repository.GetTransactions()
	if err != nil {
		util.Error(w, http.StatusInternalServerError, err)
		return
	}

	if len(transactions) == 0 {
		util.Error(w, http.StatusNotFound, fmt.Errorf("No transactions found"))
		return
	}

	util.JSON(w, http.StatusOK, transactions)
}
