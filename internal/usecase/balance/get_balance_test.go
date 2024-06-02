package balance_test

import (
	"errors"
	"testing"

	"github.com/josuedeavila/supreme-palm-tree/internal/entity"
	"github.com/josuedeavila/supreme-palm-tree/internal/entity/fake"
	"github.com/josuedeavila/supreme-palm-tree/internal/usecase/balance"
	"github.com/matryer/is"
)

func TestUseCase_New(t *testing.T) {
	is := is.New(t)
	repo := fake.NewBalanceRepository(nil, nil, nil)
	useCase := balance.New(repo)
	is.True(useCase != nil)
}

func TestUseCase_Get(t *testing.T) {
	t.Run("balance not found", func(t *testing.T) {
		is := is.New(t)
		expectedErr := errors.New("balance not found")
		repo := fake.NewBalanceRepository(nil, func(accountID int) (*entity.Balance, error) {
			return nil, expectedErr
		}, nil)
		is.True(repo != nil)

		useCase := balance.New(repo)
		is.True(useCase != nil)

		balance, err := useCase.Get(1)
		is.True(balance == nil)
		is.Equal(err.Error(), expectedErr.Error()) // Fix: Call the Error() method on the error object
	})

	t.Run("balance found", func(t *testing.T) {
		is := is.New(t)
		repo := fake.NewBalanceRepository(nil, func(accountID int) (*entity.Balance, error) {
			return &entity.Balance{
				Amount:    100,
				AccountID: accountID,
			}, nil
		}, nil)
		is.True(repo != nil)

		useCase := balance.New(repo)
		is.True(useCase != nil)

		balance, err := useCase.Get(1)
		is.NoErr(err)
		is.Equal(balance.Amount, 100)
	})
}
