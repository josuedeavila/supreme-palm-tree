package fake

import (
	"github.com/josuedeavila/supreme-palm-tree/internal/usecase/balance"
)

type UseCases struct {
	GetFunc func(accountID int) (*balance.Balance, error)
}

func (u *UseCases) Get(accountID int) (*balance.Balance, error) {
	return u.GetFunc(accountID)
}

func NewUseCases(getFunc func(accountID int) (*balance.Balance, error)) *UseCases {
	return &UseCases{
		GetFunc: getFunc,
	}
}
