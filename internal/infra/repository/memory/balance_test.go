package memory_test

import (
	"strings"
	"testing"

	repository "github.com/josuedeavila/supreme-palm-tree/internal/infra/repository/memory"

	"github.com/matryer/is"
)

func TestBalancerRepository_NewBalanceRepository(t *testing.T) {
	is := is.New(t)
	repo := repository.NewBalanceRepository()
	is.True(repo != nil)
}

func TestBalancerRepository_Methods(t *testing.T) {
	t.Run("balance not found", func(t *testing.T) {
		is := is.New(t)
		repo := repository.NewBalanceRepository()
		balance, err := repo.Get(1)
		is.True(balance == nil)
		is.True(strings.Contains(err.Error(), "balance not found"))
	})

	t.Run("balance found", func(t *testing.T) {
		is := is.New(t)
		repo := repository.NewBalanceRepository()
		_, err := repo.Update(1, 100)
		is.NoErr(err)

		balance, err := repo.Get(1)
		is.NoErr(err)
		is.Equal(balance.Amount, 100)
	})

	t.Run("delete all balances", func(t *testing.T) {
		is := is.New(t)
		repo := repository.NewBalanceRepository()
		_, err := repo.Update(1, 100)
		is.NoErr(err)

		err = repo.DeleteAll()
		is.NoErr(err)

		balance, err := repo.Get(1)
		is.True(balance == nil)
		is.True(strings.Contains(err.Error(), "balance not found"))
	})

}
