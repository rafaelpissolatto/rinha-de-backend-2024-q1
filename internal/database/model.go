package database

// customerModel is the model for the customer entity
type customerModel struct {
	// ID is the primary key
	ID int
	// Name is the name of the customer
	Name string
	// Limit is the limit of the customer
	Limit int
	// Balance is the balance of the customer
	Balance int
}

// transactionModel is the model for the transaction entity
type transactionModel struct {
	// ID is the primary key
	ID int
	// CustomerID is the foreign key to the customer
	CustomerID int
	// Value is the value of the transaction
	Value int
	// Type is the type of the transaction
	Type string
	// Description is the description of the transaction
	Description string
	// CreatedAt is the date and time of the transaction
	CreatedAt string
}
