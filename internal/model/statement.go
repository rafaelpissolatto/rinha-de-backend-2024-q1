package model

import "time"

type StatementDate time.Time

// Statement represents an statement/extract of a customer
type StatementResponse struct {
	CustomerID         int                 `json:"-"`                  // hiden field in json
	Balance            TransactionResponse `json:"saldo"`              // Balance is the current amount of credit that the customer has (could be zero)
	LatestTransactions []Transaction       `json:"ultimas_transacoes"` // LatestTransactions is a list of the latest transactions
}
