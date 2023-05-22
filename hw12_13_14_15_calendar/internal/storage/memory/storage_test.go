package memorystorage

import (
	"testing"
	"time"

	"github.com/adrevin/ogdphw/hw12_13_14_15_calendar/internal/entities"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestStorage(t *testing.T) {
	storage := New()
	evt := entities.Event{
		Title:    "Title",
		Time:     time.Now().UTC(),
		Duration: time.Hour,
		UserID:   uuid.New(),
	}

	t.Run("create", func(t *testing.T) {
		evtID := storage.Create(&evt)
		require.NotEqual(t, uuid.Nil, evtID)
	})

	t.Run("update", func(t *testing.T) {
		err := storage.Update(uuid.New(), &entities.Event{})
		require.ErrorIs(t, ErrEventNotFound, err)

		evtID := storage.Create(&evt)
		err = storage.Update(evtID, &entities.Event{Title: "NewTitle"})
		require.NoError(t, err)
	})

	t.Run("delete", func(t *testing.T) {
		err := storage.Delete(uuid.New())
		require.ErrorIs(t, ErrEventNotFound, err)

		evtID := storage.Create(&evt)
		err = storage.Delete(evtID)
		require.NoError(t, err)
	})
}
