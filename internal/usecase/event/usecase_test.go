package event_test

import (
	"testing"

	"github.com/josuedeavila/supreme-palm-tree/internal/entity/fake"
	"github.com/josuedeavila/supreme-palm-tree/internal/usecase/event"
	"github.com/matryer/is"
)

func TestUseCase_New(t *testing.T) {
	is := is.New(t)
	eventRepo := fake.NewEventRepository(nil, nil)
	balanceRepo := fake.NewBalanceRepository(nil, nil, nil)
	useCase := event.New(eventRepo, balanceRepo)
	is.True(useCase != nil)
}
