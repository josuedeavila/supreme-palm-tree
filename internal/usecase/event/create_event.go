package event

import (
	"fmt"
	"strconv"

	"github.com/josuedeavila/supreme-palm-tree/internal/entity"
)

// Create creates a new event
func (uc *useCases) Create(input *Event) (*EventOutput, error) {
	if input == nil {
		return nil, fmt.Errorf("event is required")
	}

	input.Type.Validate()
	if input.Type == InvalidEvent {
		return nil, fmt.Errorf("invalid transaction type")
	}

	inputEvent := &entity.Event{
		Type:        input.Type.String(),
		Origin:      input.Origin,
		Amount:      input.Amount,
		Destination: input.Destination,
		ID:          input.ID,
	}
	_, err := uc.eventRepository.Create(inputEvent)
	if err != nil {
		return nil, fmt.Errorf("could not create event: %w", err)
	}
	result, err := uc.DoTransaction(inputEvent)
	if err != nil {
		return nil, fmt.Errorf("could not do transaction: %w", err)
	}

	return result, nil
}

func (uc *useCases) DoTransaction(event *entity.Event) (*EventOutput, error) {
	if event == nil {
		return nil, fmt.Errorf("event is required")
	}

	switch event.Type {
	case DepositEvent.String():
		return uc.processDeposit(event)
	case WithdrawEvent.String():
		return uc.processWithdraw(event)
	case TransferEvent.String():
		return uc.processTransfer(event)
	default:
		return nil, fmt.Errorf("invalid transaction type")
	}
}

func (uc *useCases) processDeposit(event *entity.Event) (*EventOutput, error) {
	b, err := uc.updateBalance(event.Destination, event.Amount)
	if err != nil {
		return nil, err
	}

	return &EventOutput{
		Destination: &TransactionResult{
			ID:      event.Destination,
			Balance: b,
		},
	}, nil
}

func (uc *useCases) processWithdraw(event *entity.Event) (*EventOutput, error) {
	balance, err := uc.updateBalance(event.Origin, -event.Amount)
	if err != nil {
		return nil, err
	}

	return &EventOutput{
		Origin: &TransactionResult{
			ID:      event.Origin,
			Balance: balance,
		},
	}, nil
}

func (uc *useCases) processTransfer(event *entity.Event) (*EventOutput, error) {
	originBalance, err := uc.updateBalance(event.Origin, -event.Amount)
	if err != nil {
		return nil, err
	}
	destinationBalance, err := uc.updateBalance(event.Destination, event.Amount)
	if err != nil {
		return nil, err
	}

	return &EventOutput{
		Origin: &TransactionResult{
			ID:      event.Origin,
			Balance: originBalance,
		},
		Destination: &TransactionResult{
			ID:      event.Destination,
			Balance: destinationBalance,
		},
	}, nil
}

func (uc *useCases) updateBalance(account string, amount int) (int, error) {
	accountInt, err := strconv.Atoi(account)
	if err != nil {
		return 0, fmt.Errorf("could not convert account to int: %w", err)
	}

	balance, err := uc.balaceRepository.Get(accountInt)
	if err != nil {
		if err.Error() == "balance not found" && amount < 0 {
			return 0, fmt.Errorf("could not get balance: %w", err)
		}
	}

	if balance == nil {
		balance = &entity.Balance{
			AccountID: accountInt,
			Amount:    0,
		}
	}

	balance.Amount += amount
	newBalance, err := uc.balaceRepository.Update(accountInt, balance.Amount)
	if err != nil {
		return 0, fmt.Errorf("could not update balance: %w", err)
	}
	return newBalance.Amount, nil
}
