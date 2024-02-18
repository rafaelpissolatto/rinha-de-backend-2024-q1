package services

import (
	"rinha-backend-2024q1/internal/domain/entities"
	"rinha-backend-2024q1/internal/domain/repositories"
)

type Customer struct {
	repo repositories.ICustomer
}

func NewCustomers(repo repositories.ICustomer) *Customer {
	return &Customer{repo: repo}
}

func (c *Customer) GetAll() (entities.Customer, error) {
	return c.repo.GetAll()
}

func (c *Customer) GetByID(id int) (entities.Customer, error) {
	return c.repo.GetByID(id)
}

func (c *Customer) Create(customer entities.Customer) (entities.Customer, error) {
	return c.repo.Create(customer)
}

func (c *Customer) Update(id int, customer entities.Customer) (entities.Customer, error) {
	return c.repo.Update(id, customer)
}
