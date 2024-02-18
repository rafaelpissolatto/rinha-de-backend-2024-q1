package repositories

import "rinha-backend-2024q1/internal/domain/entities"

type ITransaction interface {
	GetAll() (entities.Transaction, error)
	GetByID(id int) (entities.Transaction, error)
	Create(transaction entities.Transaction) (entities.Transaction, error)
	Update(id int, transaction entities.Transaction) (entities.Transaction, error)
}

type transaction struct {
}
