package fake

import (
	entity "github.com/josuedeavila/supreme-palm-tree/internal/entity"
)

type BalanceRepository struct {
	UpdateFunc    func(accountID, amount int) (*entity.Balance, error)
	GetFunc       func(accountID int) (*entity.Balance, error)
	DeleteAllFunc func() error
}

func (b *BalanceRepository) Update(accountID, amount int) (*entity.Balance, error) {
	return b.UpdateFunc(accountID, amount)
}

func (b *BalanceRepository) Get(accountID int) (*entity.Balance, error) {
	return b.GetFunc(accountID)
}

func (b *BalanceRepository) DeleteAll() error {
	return b.DeleteAllFunc()
}

func NewBalanceRepository(updateFunc func(accountID int, amount int) (*entity.Balance, error),
	getFunc func(accountID int) (*entity.Balance, error),
	deleteAllFunc func() error) *BalanceRepository {
	return &BalanceRepository{
		UpdateFunc:    updateFunc,
		GetFunc:       getFunc,
		DeleteAllFunc: deleteAllFunc,
	}
}
