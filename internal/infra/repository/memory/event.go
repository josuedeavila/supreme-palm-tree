package memory

import (
	"sync"

	"github.com/josuedeavila/supreme-palm-tree/internal/entity"
)

type EventRepository struct {
	data   map[int]*entity.Event
	mu     sync.RWMutex
	nextID int
}

// NewEventRepository instantiates offer repository
func NewEventRepository() entity.EventRepository {
	return &EventRepository{
		data: make(map[int]*entity.Event),
		mu:   sync.RWMutex{},
	}
}

// Create creates a new event
func (r *EventRepository) Create(event *entity.Event) (*entity.Event, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	event.ID = r.nextID
	r.data[event.ID] = event
	r.nextID++
	return event, nil
}
