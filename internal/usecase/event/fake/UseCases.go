package fake

import (
	"github.com/josuedeavila/supreme-palm-tree/internal/usecase/event"
)

type UseCases struct {
	CreateFunc func(*event.Event) (*event.EventOutput, error)
}

func (u *UseCases) Create(event *event.Event) (*event.EventOutput, error) {
	return u.CreateFunc(event)
}

func NewUseCases(createFunc func(*event.Event) (*event.EventOutput, error)) *UseCases {
	return &UseCases{
		CreateFunc: createFunc,
	}
}
