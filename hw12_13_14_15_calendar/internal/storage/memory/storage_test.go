package memorystorage

import (
	"testing"
	"time"

	"github.com/adrevin/ogdphw/hw12_13_14_15_calendar/internal/entities"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestStorage(t *testing.T) {
	t.Run("basic", func(t *testing.T) {
		evt := entities.Event{
			ID:       uuid.Nil,
			Title:    "Title",
			Time:     time.Now().UTC(),
			Duration: time.Hour,
			UserID:   uuid.New(),
		}
		storage := &Storage{}
		evtID, err := storage.Create(evt)
		require.NoError(t, err)
		require.NotEqual(t, uuid.Nil, evtID)
	})
}
