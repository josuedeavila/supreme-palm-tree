package entity

type TransactionType int

const (
	Deposit TransactionType = iota
	Withdraw
	Transfer
)

func (t TransactionType) String() string {
	return [...]string{"Deposit", "Withdraw", "Transfer"}[t]
}

// Event entity table model
type Event struct {
	ID          int             `json:"id"`
	Type        TransactionType `json:"type"`
	Origin      string          `json:"origin"`
	Amount      int             `json:"amount"`
	Destination string          `json:"destination"`
}

// EventRepository is the interface for the Event repository
type EventRepository interface {
	Create(event *Event) (*Event, error)
}
