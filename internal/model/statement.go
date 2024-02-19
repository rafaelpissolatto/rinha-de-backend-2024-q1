package model

import "time"

type StatementDate time.Time

// Statement represents an statement/extract of a customer
type Statement struct {
	CustomerID uint `json:"-"` // hiden field in json
	Balance    struct {
		Total int `json:"total"` //Total will return the current balance of the customer
		//StatementDate will return the current date of the statement ("2024-01-17T02:34:41.217753Z")
		StatementDate StatementDate `json:"data"`
		Credit_Limit  Credit_Limit  `json:"limite"` //Credit_Limit will return the credit limit of the customer
	} `json:"saldo"`
	LatestTransactions []Transaction `json:"ultimas_transacoes"` // LatestTransactions is a list of the latest transactions
}
