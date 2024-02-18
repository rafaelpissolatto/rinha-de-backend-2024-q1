package entities

type Transaction struct {
	ID          int
	CustomerID  int
	Value       int
	Type        string
	Description string
	CreatedAt   string
}

type Transactions []Transaction
