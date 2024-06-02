package memory_test

import (
	"testing"

	"github.com/josuedeavila/supreme-palm-tree/internal/entity"
	repository "github.com/josuedeavila/supreme-palm-tree/internal/infra/repository/memory"

	"github.com/matryer/is"
)

func TestEventRepository_NewEventRepository(t *testing.T) {
	is := is.New(t)
	repo := repository.NewEventRepository()
	is.True(repo != nil)
}

func TestEventRepository_Methods(t *testing.T) {
	t.Run("event created", func(t *testing.T) {
		is := is.New(t)
		repo := repository.NewEventRepository()
		event := &entity.Event{}
		event, err := repo.Create(event)
		is.NoErr(err)
		is.True(event.ID == 0)
	})

	t.Run("delete all events", func(t *testing.T) {
		is := is.New(t)
		repo := repository.NewEventRepository()
		event := &entity.Event{}
		_, err := repo.Create(event)
		is.NoErr(err)
		err = repo.DeleteAll()
		is.NoErr(err)
	})
}
