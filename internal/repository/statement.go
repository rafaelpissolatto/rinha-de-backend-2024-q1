package repository

import (
	"errors"
	"log"
	"rinha-backend-2024q1-rafael-pissolatto-nunes/internal/model"
)

// GetSimpleStatementByCustomerId returns current balance and credit limit for a customer
func (c *Customers) GetSimpleStatementByCustomerId(customerID int) (model.TransactionResponse, error) {
	var customer model.Customer
	err := c.db.QueryRow("SELECT id, credit_limit, balance FROM customers WHERE id = ?", customerID).Scan(&customer.ID, &customer.CreditLimit, &customer.Balance)
	if err != nil {
		return model.TransactionResponse{}, err
	}

	return model.TransactionResponse{Limit: customer.CreditLimit, Balance: customer.Balance}, nil
}

// GetStatement returns a list of transactions for a customer
func (c *Customers) GetCompleteStatementByCustomerId(customerID int) (model.StatementResponse, error) {
	rows, err := c.db.Query("SELECT id, customer_id, amount, operator_type, description, created_at FROM transactions WHERE customer_id = ? ORDER BY created_at", customerID)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return model.StatementResponse{}, errors.New("No transactions found")
		}
		log.Println("[ERROR] Failed to get the transactions: [", err, "]")
		return model.StatementResponse{}, err
	}
	defer rows.Close()

	transactions := []model.Transaction{}
	for rows.Next() {
		var transaction model.Transaction
		if err := rows.Scan(&transaction.ID, &transaction.CustomerID, &transaction.Amount, &transaction.OperatorType, &transaction.Description, &transaction.CreatedAt); err != nil {
			log.Println("[ERROR] Failed to scan the transaction: [", err, "]")
			return model.StatementResponse{}, err
		}
		transactions = append(transactions, transaction)
	}

	balance := model.TransactionResponse{}
	err = c.db.QueryRow("SELECT credit_limit, balance FROM customers WHERE id = ?", customerID).Scan(&balance.Limit, &balance.Balance)
	if err != nil {
		log.Println("[ERROR] Failed to get the customer balance: [", err, "]")
		return model.StatementResponse{}, err
	}

	return model.StatementResponse{
		Balance:            balance,
		LatestTransactions: transactions,
	}, nil
}
