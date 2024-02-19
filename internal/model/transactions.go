package model

import "errors"

// Transaction represents a transaction
type Transaction struct {
	ID           uint   `json:"-"` // hiden field in json
	CustomerID   uint   `json:"-"` // hiden field in json
	Amount       int    `json:"valor"`
	OperatorType string `json:"tipo"`
	Description  string `json:"descricao"`
	CreatedAt    string `json:"criado_em"`
}

// validate validates the transaction
func (t *Transaction) validate() error {
	switch OperatorType {
	case "c": //credit
		if t.CustomerID == 0 {
			return errors.New("customer_id is required")
		}
		if t.Amount == 0 {
			return errors.New("amount is required")
		}
		if t.OperatorType == "" {
			return errors.New("operator_type is required")
		}
		if t.Description == "" {
			return errors.New("description is required")
		}
		if t.Amount < 

	case "d": //debit
		if t.CustomerID == 0 {
			return errors.New("customer_id is required")
		}
		if t.Amount == 0 {
			return errors.New("amount is required")
		}
		if t.OperatorType == "" {
			return errors.New("operator_type is required")
		}
		if t.Description == "" {
			return errors.New("description is required")
		}
	}

	return nil
}

// Prepare prepares the transaction to be saved
func (t *Transaction) Prepare(step string) error {
	if err := t.validate(step); err != nil {
		return err
	}

	return nil
}
