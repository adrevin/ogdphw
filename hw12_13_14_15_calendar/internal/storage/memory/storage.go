package memorystorage

import (
	"sync"
	"time"

	"github.com/adrevin/ogdphw/hw12_13_14_15_calendar/internal/entities"
	"github.com/google/uuid"
)

type Storage struct {
	mu sync.RWMutex //nolint:unused
}

func New() *Storage {
	return &Storage{}
}

func (l Storage) Create(event entities.Event) (uuid.UUID, error) { //nolint:govet
	event.ID = uuid.New()
	// TODO
	return event.ID, nil
}

func (l Storage) Update(id uuid.UUID, event entities.Event) error { //nolint:govet,revive
	// TODO
	return nil
}

func (l Storage) Delete(id uuid.UUID) error { //nolint:govet,revive
	// TODO
	return nil
}

func (l Storage) DayEvens(time time.Time) ([]entities.Event, error) { //nolint:govet,revive
	// TODO
	return make([]entities.Event, 0), nil
}

func (l Storage) WeekEvens(time time.Time) ([]entities.Event, error) { //nolint:govet,revive
	// TODO
	return make([]entities.Event, 0), nil
}

func (l Storage) MonthEvens(time time.Time) ([]entities.Event, error) { //nolint:govet,revive
	// TODO
	return make([]entities.Event, 0), nil
}
