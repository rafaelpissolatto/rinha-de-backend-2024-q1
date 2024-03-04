package repository

import (
	"database/sql"
	"errors"
	"log"
	"rinha-backend-2024q1-rafael-pissolatto-nunes/internal/model"
)

// Represents a repository for Customers
type Customers struct {
	db *sql.DB
}

// NewCustomers creates a new Customers repository
func NewCustomersRepository(db *sql.DB) *Customers {
	return &Customers{db}
}

// GetCustomers returns a list of customers (for now, all customers)
func (c *Customers) GetCustomers() ([]model.Customer, error) {
	rows, err := c.db.Query("SELECT id, credit_limit, balance FROM customers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	customers := []model.Customer{}
	for rows.Next() {
		var customer model.Customer
		if err := rows.Scan(&customer.ID, &customer.CreditLimit, &customer.Balance); err != nil {
			return nil, err
		}
		customers = append(customers, customer)
	}

	return customers, nil
}

// GetCustomer returns a customer by its ID
func (c *Customers) GetCustomerByID(id int) (*model.Customer, error) {
	var customer model.Customer
	err := c.db.QueryRow("SELECT id, credit_limit, balance FROM customers WHERE id = ?", id).Scan(&customer.ID, &customer.CreditLimit, &customer.Balance)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, errors.New("Customer not found")
		}
		log.Println("[ERROR] Failed to get the customer: [", err, "]")
		return nil, err
	}

	return &customer, nil
}
