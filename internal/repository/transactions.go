package repository

import (
	"database/sql"
	"errors"
	"log"
	"rinha-backend-2024q1-rafael-pissolatto-nunes/internal/model"
	"time"
)

// Represents a repository for Transactions
type Transactions struct {
	db *sql.DB
}

// NewTransactions creates a new Transactions repository
func NewTransactionsRepository(db *sql.DB) *Transactions {
	return &Transactions{db}
}

// GetTransactions returns a list of transactions (for now, all transactions)
func (t *Transactions) GetTransactions() ([]model.Transaction, error) {
	rows, err := t.db.Query("SELECT id, customer_id, amount, operator_type, description, created_at FROM transactions")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	transactions := []model.Transaction{}
	for rows.Next() {
		var transaction model.Transaction
		if err := rows.Scan(&transaction.ID, &transaction.CustomerID, &transaction.Amount, &transaction.OperatorType, &transaction.Description, &transaction.CreatedAt); err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}

	return transactions, nil
}

// GetTransaction returns a transaction by its ID
func (t *Transactions) GetTransaction(id int) (*model.Transaction, error) {
	var transaction model.Transaction
	err := t.db.QueryRow("SELECT id, customer_id, amount, operator_type, description, created_at FROM transactions WHERE id = ?", id).Scan(&transaction.ID, &transaction.CustomerID, &transaction.Amount, &transaction.OperatorType, &transaction.Description, &transaction.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &transaction, nil
}

// GetTransactionsByCustomerId returns a list of transactions for a customer
func (t *Transactions) GetTransactionsByCustomerId(id int) ([]model.Transaction, error) {
	rows, err := t.db.Query("SELECT id, customer_id, amount, operator_type, description, created_at FROM transactions WHERE customer_id = ? ORDER BY created_at", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	transactions := []model.Transaction{}
	for rows.Next() {
		var transaction model.Transaction
		if err := rows.Scan(&transaction.ID, &transaction.CustomerID, &transaction.Amount, &transaction.OperatorType, &transaction.Description, &transaction.CreatedAt); err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}

	return transactions, nil
}

// CreateTransaction creates a new transaction for a customer
// this function should be atomic and should not allow the customer's balance to be less than their available limit
// Also, this function should return an error if the customer does not exist (404)
// And, this function should return an error if the transaction is not valid (400)
// Should be atomic and should handle many concurrent transactions
func (c *Customers) CreateTransaction(customerID int, transaction model.Transaction) error {
	// Start a transaction
	tx, err := c.db.Begin()
	if err != nil {
		return err
	}

	// Check if the customer exists
	var customer model.Customer
	err = tx.QueryRow("SELECT id, credit_limit, balance FROM customers WHERE id = ?", customerID).Scan(&customer.ID, &customer.CreditLimit, &customer.Balance)
	if err != nil {
		log.Println("[ERROR] Failed to execute the statement: [", err, "]")
		tx.Rollback() // Call the Rollback method
		return err
	}
	log.Println("[TRACE] Customer: ", customer)

	// Check if the transaction is valid
	err = transaction.Prepare()
	if err != nil {
		log.Println("[ERROR] Failed to execute the statement: [", err, "]")
		tx.Rollback() // Call the Rollback method
		return err
	}

	// Check if the customer's balance will be less than their available limit
	switch transaction.OperatorType {
	case "c": //credit
		if customer.Balance+transaction.Amount > customer.CreditLimit {
			log.Println("[DEBUG] customerCreditLimit: ", customer.CreditLimit, "customerBalance+transactionAmount: ", customer.Balance+transaction.Amount)
			tx.Rollback() // Call the Rollback method
			return errors.New("The customer's balance will be less than their available limit")
		}
	case "d": //debit
		if customer.Balance-transaction.Amount < -customer.CreditLimit {
			log.Println("[DEBUG] customerCreditLimit: ", customer.CreditLimit, "customerBalance-transactionAmount: ", customer.Balance-transaction.Amount)
			tx.Rollback() // Call the Rollback method
			return errors.New("The customer's balance will be less than their available limit")
		}
	default:
		log.Println("[ERROR] Invalid operator type")
		tx.Rollback() // Call the Rollback method
		return errors.New("Invalid operator type")
	}

	// Create the transaction
	createAt := time.Now().Format("2006-01-02 15:04:05")
	log.Println("[TRACE] Transaction: ", transaction)

	_, err = tx.Exec(
		"INSERT INTO transactions (customer_id, amount, operator_type, description, created_at) VALUES (?, ?, ?, ?, ?)", transaction.CustomerID, transaction.Amount, transaction.OperatorType, transaction.Description, createAt)
	if err != nil {
		tx.Rollback() // Call the Rollback method
		return err
	}

	// Update the customer's balance
	switch transaction.OperatorType {
	case "c": //credit
		customer.Balance += transaction.Amount
	case "d": //debit
		customer.Balance -= transaction.Amount
	default:
		tx.Rollback()
		return errors.New("Invalid operator type")
	}

	_, err = tx.Exec("UPDATE customers SET balance = ? WHERE id = ?", customer.Balance, transaction.CustomerID)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

// PostTransactionByCustomerId creates a new transaction for a customer
// this function should be atomic and should not allow the customer's balance to be less than their available limit
// Also, this function should return an error if the customer does not exist (404)
// And, this function should return an error if the transaction is not valid (400)
// Should be atomic and should handle many concurrent transactions
func (c *Customers) PostTransactionByCustomerId(id int, transaction model.Transaction) (model.Transaction, error) {
	// Start a transaction
	tx, err := c.db.Begin()
	if err != nil {
		return transaction, err
	}

	// Check if the customer exists
	var customer model.Customer
	err = tx.QueryRow("SELECT id, credit_limit, balance FROM customers WHERE id = ?", id).Scan(&customer.ID, &customer.CreditLimit, &customer.Balance)
	if err != nil {
		log.Println("[ERROR] Failed to execute the statement2: [", err, "]")
		tx.Rollback() // Call the Rollback method
		return transaction, err
	}
	log.Println("[TRACE] Customer2: ", customer)

	// Check if the transaction is valid
	err = transaction.Prepare()
	if err != nil {
		log.Println("[ERROR] Failed to execute the statement3: [", err, "]")
		tx.Rollback() // Call the Rollback method
		return transaction, err
	}

	// Check if the customer's balance will be less than their available limit
	switch transaction.OperatorType {
	case "c": //credit
		if customer.Balance+transaction.Amount > customer.CreditLimit {
			log.Println("[DEBUG] customerCreditLimit: ", customer.CreditLimit, "customerBalance+transactionAmount: ", customer.Balance+transaction.Amount)
			log.Println("[WARN] The customer's balance will be less than their available limit")
			tx.Rollback() // Call the Rollback method
			return transaction, err
		}
	case "d": //debit
		if customer.Balance-transaction.Amount < -customer.CreditLimit {
			log.Println("[DEBUG] customerCreditLimit: ", customer.CreditLimit, "customerBalance-transactionAmount: ", customer.Balance-transaction.Amount)
			log.Panicln("[WARN] The customer's balance will be less than their available limit")
			tx.Rollback() // Call the Rollback method
			return transaction, err
		}
	default:
		log.Panicln("[WARN] The customer's balance will be less than their available limit")
		tx.Rollback() // Call the Rollback method
		return transaction, err
	}
	log.Println(transaction)

	// Create the transaction
	_, err = tx.Exec("INSERT INTO transactions (customer_id, amount, created_at) VALUES (?, ?, ?)", transaction.CustomerID, transaction.Amount, transaction.CreatedAt)
	if err != nil {
		tx.Rollback() // Call the Rollback method
		log.Println("[ERROR] Failed to execute the statement3: [", err, "]")
		return transaction, err
	}

	// Update the customer's balance
	switch transaction.OperatorType {
	case "c": //credit
		customer.Balance += transaction.Amount
	case "d": //debit
		customer.Balance -= transaction.Amount
	default:
		log.Println("[ERROR] Invalid operator type")
		tx.Rollback()
		return transaction, err
	}

	_, err = tx.Exec("UPDATE customers SET balance = ? WHERE id = ?", customer.Balance, transaction.CustomerID)
	if err != nil {
		log.Println("[ERROR] Failed to execute the statement4: [", err, "]")
		tx.Rollback()
		return transaction, err
	}

	err = tx.Commit()
	if err != nil {
		log.Println("[ERROR] Failed to execute the statement5: [", err, "]")
		return transaction, err
	}

	log.Println("[INFO] Transaction created: ", transaction)
	return transaction, nil
}
