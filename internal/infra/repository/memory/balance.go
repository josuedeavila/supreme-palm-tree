package memory

import (
	"fmt"
	"sync"

	"github.com/josuedeavila/supreme-palm-tree/internal/entity"
)

type BalanceRepository struct {
	data map[int]*entity.Balance
	mu   sync.RWMutex
}

// NewBalanceRepository instantiates offer repository
func NewBalanceRepository() entity.BalanceRepository {
	return &BalanceRepository{
		data: make(map[int]*entity.Balance),
		mu:   sync.RWMutex{},
	}
}

// Get retrieves a balance by accountID
func (r *BalanceRepository) Get(accountID int) (*entity.Balance, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	balance, ok := r.data[accountID]
	if !ok {
		return nil, fmt.Errorf("balance not found")
	}
	return balance, nil
}

// Update updates a balance by accountID
func (r *BalanceRepository) Update(accountID int, amount int) (*entity.Balance, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	balance, ok := r.data[accountID]
	if !ok {
		balance = &entity.Balance{
			AccountID: accountID,
		}
		r.data[accountID] = balance
	}
	balance.Amount = amount
	return balance, nil
}

// DeleteAll deletes all balances
func (r *BalanceRepository) DeleteAll() error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.data = make(map[int]*entity.Balance)
	return nil
}
