package dto

// Balance entity table model
type Balance struct {
	Amount    int `json:"amount"`
	AccountID int `json:"account_id"`
}
