package fake

import (
	entity "github.com/josuedeavila/supreme-palm-tree/internal/entity"
)

type EventRepository struct {
	CreateFunc    func(event *entity.Event) (*entity.Event, error)
	DeleteAllFunc func() error
}

func (r *EventRepository) Create(event *entity.Event) (*entity.Event, error) {
	return r.CreateFunc(event)
}

func (r *EventRepository) DeleteAll() error {
	return r.DeleteAllFunc()
}

func NewEventRepository(createFunc func(event *entity.Event) (*entity.Event, error),
	deleterAllFunc func() error) *EventRepository {
	return &EventRepository{
		CreateFunc:    createFunc,
		DeleteAllFunc: deleterAllFunc,
	}
}
