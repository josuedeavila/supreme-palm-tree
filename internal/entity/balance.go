package entity

// Balance entity table model
type Balance struct {
	Amount    int `json:"amount"`
	AccountID int `json:"account_id"`
}

// BalanceRepository is the interface for the Balance repository
type BalanceRepository interface {
	Get(accountID int) (*Balance, error)
	Update(accountID int, amount int) (*Balance, error)
}
