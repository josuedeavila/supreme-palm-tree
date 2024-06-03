package event_test

import (
	"fmt"
	"testing"

	"github.com/josuedeavila/supreme-palm-tree/internal/entity"
	"github.com/josuedeavila/supreme-palm-tree/internal/entity/fake"
	"github.com/josuedeavila/supreme-palm-tree/internal/usecase/event"
	"github.com/matryer/is"
)

func TestUseCase_Create(t *testing.T) {
	t.Run("event is required", func(t *testing.T) {
		is := is.New(t)
		useCase := event.New(nil, nil)
		is.True(useCase != nil)

		event, err := useCase.Create(nil)
		is.True(event == nil)
		is.Equal(err, fmt.Errorf("event is required"))
	})

	t.Run("ivalid event type", func(t *testing.T) {
		is := is.New(t)
		useCase := event.New(nil, nil)
		is.True(useCase != nil)

		input := &event.Event{
			Type: "",
		}
		event, err := useCase.Create(input)
		is.True(event == nil)
		is.Equal(err, fmt.Errorf("invalid transaction type"))
	})

	t.Run("event not created", func(t *testing.T) {
		is := is.New(t)
		expectedErr := fmt.Errorf("error")
		eventRepo := fake.NewEventRepository(func(event *entity.Event) (*entity.Event, error) {
			return nil, expectedErr
		}, nil)
		is.True(eventRepo != nil)

		useCase := event.New(eventRepo, nil)
		is.True(useCase != nil)

		input := &event.Event{
			Type:        "deposit",
			Origin:      "origin",
			Amount:      100,
			Destination: "destination",
		}
		event, err := useCase.Create(input)
		is.True(event == nil)
		is.Equal(err, fmt.Errorf("could not create event: %w", expectedErr))
	})

	t.Run("deposit event created", func(t *testing.T) {
		is := is.New(t)
		eventRepo := fake.NewEventRepository(func(event *entity.Event) (*entity.Event, error) {
			return &entity.Event{
				Type:        event.Type,
				Origin:      event.Origin,
				Amount:      event.Amount,
				Destination: event.Destination,
			}, nil
		}, nil)
		is.True(eventRepo != nil)

		balanceRepo := fake.NewBalanceRepository(func(accountID, amount int) (*entity.Balance, error) {
			return &entity.Balance{
				AccountID: accountID,
				Amount:    amount,
			}, nil
		}, func(accountID int) (*entity.Balance, error) {
			return &entity.Balance{
				AccountID: accountID,
				Amount:    100,
			}, nil
		}, nil)

		useCase := event.New(eventRepo, balanceRepo)
		is.True(useCase != nil)

		input := &event.Event{
			Type:        "deposit",
			Origin:      "",
			Amount:      100,
			Destination: "1",
		}
		event, err := useCase.Create(input)
		is.NoErr(err)
		is.True(event.Origin == nil)
		is.True(event.Destination != nil)
	})

	t.Run("withdraw event created", func(t *testing.T) {
		is := is.New(t)
		eventRepo := fake.NewEventRepository(func(event *entity.Event) (*entity.Event, error) {
			return &entity.Event{
				Type:        event.Type,
				Origin:      event.Origin,
				Amount:      event.Amount,
				Destination: event.Destination,
			}, nil
		}, nil)
		is.True(eventRepo != nil)

		balanceRepo := fake.NewBalanceRepository(func(accountID, amount int) (*entity.Balance, error) {
			return &entity.Balance{
				AccountID: accountID,
				Amount:    amount,
			}, nil
		}, func(accountID int) (*entity.Balance, error) {
			return &entity.Balance{
				AccountID: accountID,
				Amount:    100,
			}, nil
		}, nil)

		useCase := event.New(eventRepo, balanceRepo)
		is.True(useCase != nil)

		input := &event.Event{
			Type:        "withdraw",
			Origin:      "1",
			Amount:      100,
			Destination: "",
		}
		event, err := useCase.Create(input)
		is.NoErr(err)
		is.True(event.Origin != nil)
		is.True(event.Destination == nil)
	})

	t.Run("transfer event created", func(t *testing.T) {
		is := is.New(t)
		eventRepo := fake.NewEventRepository(func(event *entity.Event) (*entity.Event, error) {
			return &entity.Event{
				Type:        event.Type,
				Origin:      event.Origin,
				Amount:      event.Amount,
				Destination: event.Destination,
			}, nil
		}, nil)
		is.True(eventRepo != nil)

		balanceRepo := fake.NewBalanceRepository(func(accountID, amount int) (*entity.Balance, error) {
			return &entity.Balance{
				AccountID: accountID,
				Amount:    amount,
			}, nil
		}, func(accountID int) (*entity.Balance, error) {
			return &entity.Balance{
				AccountID: accountID,
				Amount:    100,
			}, nil
		}, nil)

		useCase := event.New(eventRepo, balanceRepo)
		is.True(useCase != nil)

		input := &event.Event{
			Type:        "transfer",
			Origin:      "1",
			Amount:      100,
			Destination: "2",
		}
		event, err := useCase.Create(input)
		is.NoErr(err)
		is.True(event.Origin != nil)
		is.True(event.Destination != nil)
	})
}
