package services

import (
	"rinha-backend-2024q1/internal/domain/entities"
	"rinha-backend-2024q1/internal/domain/repositories"
)

type TransactionConfiguration func(t *Transaction) error

type Transaction struct {
	repo repositories.ITransaction
}

func NewTransactions(repo repositories.ITransaction) *Transaction {
	return &Transaction{repo: repo}
}

func (t *Transaction) GetAll() (entities.Transaction, error) {
	return t.repo.GetAll()
}

func (t *Transaction) GetByID(id int) (entities.Transaction, error) {
	return t.repo.GetByID(id)
}

func (t *Transaction) Create(transaction entities.Transaction) (entities.Transaction, error) {
	return t.repo.Create(transaction)
}

func (t *Transaction) Update(id int, transaction entities.Transaction) (entities.Transaction, error) {
	return t.repo.Update(id, transaction)
}
