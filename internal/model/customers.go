package model

import (
	"errors"
)

// Customer is a struct that represents a customer
type Customer struct {
	// ID is the unique identifier of the customer
	ID uint `json:"id"`
	// Name is the name of the customer
	Name string `json:"-"`
	// Limit is the maximum amount of credit that the customer can have (could be zero)
	Credit_Limit int `json:"limite"`
	// Balance is the current amount of credit that the customer has (could be zero)
	Balance int `json:"saldo"`
}

// validate validates the customer
func (c *Customer) validate() error {
	// Rules: A debit transaction can never leave the customer's balance less than their available limit.
	// For example, a customer with a limit of 1000 (R$ 10) should never have a balance lower than -1000 (R$ -10).
	// In this case, a balance of -1001 or lower means inconsistency
	if c.Credit_Limit < 0 {
		return errors.New("limit must be greater than or equal to zero")
	}
	if c.Balance < 0 {
		return errors.New("balance must be greater than or equal to zero")
	}

	if c.Balance < -c.Credit_Limit {
		return errors.New("balance must be greater than or equal to the negative limit")
	}

	return nil
}

// Prepare prepares the customer to be saved
func (c *Customer) Prepare() error {
	if err := c.validate(); err != nil {
		return err
	}

	return nil
}
