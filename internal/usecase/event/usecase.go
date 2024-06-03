package event

import (
	"github.com/josuedeavila/supreme-palm-tree/internal/entity"
)

// TransactionType is the type of transaction
type TransactionType string

const (
	DepositEvent  TransactionType = "deposit"
	WithdrawEvent TransactionType = "withdraw"
	TransferEvent TransactionType = "transfer"
	InvalidEvent  TransactionType = "invalid"
)

// String returns the string value of the TransactionType
func (t TransactionType) String() string {
	return string(t)
}

func (t *TransactionType) Validate() *TransactionType {
	switch *t {
	case DepositEvent, WithdrawEvent, TransferEvent:
		return t
	}
	*t = InvalidEvent
	return t
}

// Event model
type Event struct {
	ID          int             `json:"id"`
	Type        TransactionType `json:"type"`
	Origin      string          `json:"origin"`
	Amount      int             `json:"amount"`
	Destination string          `json:"destination"`
}

// EventOutput model
type EventOutput struct {
	Origin      *TransactionResult `json:"origin"`
	Destination *TransactionResult `json:"destinarion"`
}

type TransactionResult struct {
	ID      string `json:"id"`
	Balance int    `json:"balance"`
}

// UseCases exposes use cases for offer
type UseCases interface {
	Create(input *Event) (*EventOutput, error)
}

type useCases struct {
	eventRepository  entity.EventRepository
	balaceRepository entity.BalanceRepository
}

// New the service responsible for all use cases for event
func New(eventRepo entity.EventRepository, balanceRepo entity.BalanceRepository) UseCases {
	return &useCases{
		eventRepository:  eventRepo,
		balaceRepository: balanceRepo,
	}
}
