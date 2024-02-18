package entities

type Customer struct {
	ID      int
	Name    string
	Limit   int
	Balance int
}

type Customers []Customer
