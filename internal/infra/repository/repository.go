package repository

import (
	"github.com/josuedeavila/supreme-palm-tree/internal/entity"
	"github.com/josuedeavila/supreme-palm-tree/internal/infra/repository/memory"
)

// Repository is a service to abstract postgresql layer
type Repository struct {
	entity.EventRepository
	entity.BalanceRepository
}

// New creates a new repository
func New() *Repository {
	return &Repository{
		EventRepository:   memory.NewEventRepository(),
		BalanceRepository: memory.NewBalanceRepository(),
	}
}
