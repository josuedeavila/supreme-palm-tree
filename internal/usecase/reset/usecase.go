package reset

import (
	"github.com/josuedeavila/supreme-palm-tree/internal/entity"
)

// UseCases exposes use cases for offer
type UseCases interface {
	Reset() error
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
