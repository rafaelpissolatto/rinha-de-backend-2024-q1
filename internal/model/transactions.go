package model

import (
	"errors"
	"log"
)

// Transaction represents a transaction
type Transaction struct {
	ID           int    `json:"-"` // hiden field in json
	CustomerID   int    `json:"-"` // hiden field in json
	Amount       int    `json:"valor"`
	OperatorType string `json:"tipo"`
	Description  string `json:"descricao"`
	CreatedAt    string `json:"realizada_em"`
}

// TransactionResponse represents a transaction response
type TransactionResponse struct {
	Limit   int `json:"limite"`
	Balance int `json:"saldo"`
}

// validate validates the transaction according to the operator type
func (t *Transaction) validate() error {
	switch t.OperatorType {
	case "c": //credit
		if t.Amount < 0 {
			return errors.New("amount must be greater than or equal to zero")
		}
	case "d": //debit
		if t.Amount < 0 {
			return errors.New("amount must be greater than or equal to zero")
		}
	default:
		log.Println("[WARN] Invalid operator type")
		return errors.New("invalid operator type")
	}

	if t.Description == "" || len(t.Description) > 10 {
		log.Println("[WARN] Description cannot be empty or greater than 10 characters")
		return errors.New("description cannot be empty or greater than 10 characters")
	}

	if t.CustomerID == 0 {
		log.Println("[WARN] Customer ID cannot be empty")
		return errors.New("customer id cannot be empty")
	}

	return nil
}

// Prepare prepares the transaction to be saved
func (t *Transaction) Prepare() error {
	if err := t.validate(); err != nil {
		return err
	}

	return nil
}
