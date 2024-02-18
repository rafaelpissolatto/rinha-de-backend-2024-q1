package repositories

import (
	"rinha-backend-2024q1/internal/domain/entities"
)

type ICustomer interface {
	GetAll() (entities.Customer, error)
	GetByID(id int) (entities.Customer, error)
	Create(customer entities.Customer) (entities.Customer, error)
	Update(id int, customer entities.Customer) (entities.Customer, error)
}

type customer struct {
}

func (c *customer) GetAll() ([]entities.Customer, error) {
	statement, err := db.Prepare("SELECT id, name, limit, balance FROM customers")
	if err != nil {
		return []entities.Customer{}, err
	}
	defer statement.Close()

	rows, err := statement.Query()
	if err != nil {
		return []entities.Customer{}, err
	}
	defer rows.Close()

	var customers []entities.Customer
	for rows.Next() {
		var customer entities.Customer
		err = rows.Scan(&customer.ID, &customer.Name, &customer.Limit, &customer.Balance)
		if err != nil {
			return []entities.Customer{}, err
		}
		customers = append(customers, customer)
	}
	return customers, nil
}

func (c *customer) GetByID(id int) (entities.Customer, error) {
	panic("implement me")
}

func (c *customer) Create(customer entities.Customer) (entities.Customer, error) {
	panic("implement me")
}

func (c *customer) Update(id int, customer entities.Customer) (entities.Customer, error) {
	panic("implement me")
}

func NewCustomer() ICustomer {
	return &customer{}
}
