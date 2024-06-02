package balance

import "github.com/josuedeavila/supreme-palm-tree/internal/entity"

// Balance entity table model
type Balance struct {
	Amount    int `json:"amount"`
	AccountID int `json:"account_id"`
}

// UseCases exposes use cases for offer
type UseCases interface {
	Get(accountID int) (*Balance, error)
	// Update(accountID int, amount int) (*Balance, error)
}

type useCases struct {
	repository entity.BalanceRepository
}

// New the service responsible for all use cases for balance
func New(repository entity.BalanceRepository) UseCases {
	return &useCases{
		repository: repository,
	}
}
