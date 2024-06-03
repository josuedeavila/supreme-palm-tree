package dto

// Event entity table model
type Event struct {
	ID          int    `json:"id"`
	Type        string `json:"type"`
	Origin      string `json:"origin"`
	Amount      int    `json:"amount"`
	Destination string `json:"destination"`
}

// EventOutput model
type EventOutput struct {
	Origin      *TransactionResult `json:"origin,omitempty"`
	Destination *TransactionResult `json:"destination,omitempty"`
}

// TransactionResult model
type TransactionResult struct {
	ID      string `json:"id"`
	Balance int    `json:"balance"`
}
