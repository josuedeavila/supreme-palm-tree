package balance

// Get gets the balance of an account
func (uc *useCases) Get(accountID int) (*Balance, error) {
	balance, err := uc.repository.Get(accountID)
	if err != nil {
		return nil, err
	}
	return &Balance{
		Amount:    balance.Amount,
		AccountID: balance.AccountID,
	}, nil
}
