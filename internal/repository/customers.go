package repository

import (
	"database/sql"
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
	rows, err := c.db.Query("SELECT id, name, credit_limit, balance FROM customers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	customers := []model.Customer{}
	for rows.Next() {
		var customer model.Customer
		if err := rows.Scan(&customer.ID, &customer.Name, &customer.Credit_Limit, &customer.Balance); err != nil {
			return nil, err
		}
		customers = append(customers, customer)
	}

	return customers, nil
}

// GetCustomer returns a customer by its ID
func (c *Customers) GetCustomer(id int) (*model.Customer, error) {
	var customer model.Customer
	err := c.db.QueryRow("SELECT id, name, credit_limit, balance FROM customers WHERE id = ?", id).Scan(&customer.ID, &customer.Name, &customer.Credit_Limit, &customer.Balance)
	if err != nil {
		return nil, err
	}

	return &customer, nil
}

// GetStatement returns a list of transactions for a customer
func (c *Customers) GetStatementByCustomerId(id int) (model.Statement, error) {
	customer := model.Customer{}
	err := c.db.QueryRow("SELECT id, name, credit_limit, balance FROM customers WHERE id = ?", id).Scan(&customer.ID, &customer.Name, &customer.Credit_Limit, &customer.Balance)
	if err != nil {
		return model.Statement{}, err
	}

	rows, err := c.db.Query("SELECT id, customer_id, amount FROM transactions WHERE customer_id = ? ORDER BY created_at", id)
	if err != nil {
		return model.Statement{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var transaction model.Transaction

		if err := rows.Scan(&transaction.ID, &transaction.CustomerID, &transaction.Amount, &transaction.Description, &transaction.CreatedAt); err != nil {
			return model.Statement{}, err
		}

		switch transaction.OperatorType {
		case "c": //credit
			customer.Balance += transaction.Amount
		case "d": //debit
			customer.Balance -= transaction.Amount
		}

		customer.Statement = append(customer.Statement, transaction)
	}

	return customer.Statement, nil
}

// CreateTransaction creates a new customer
func (c *Customers) CreateTransaction(transaction *model.Transaction) error {
	_, err := c.db.Exec("INSERT INTO transactions (customer_id, amount) VALUES (?, ?)", transaction.CustomerID, transaction.Amount)
	return err
}
