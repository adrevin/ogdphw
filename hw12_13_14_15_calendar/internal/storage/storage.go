package storage

import (
	"time"

	"github.com/google/uuid"
)

type Storage interface {
	Create(event *Event) (uuid.UUID, error)
	Update(id uuid.UUID, event *Event) error
	Delete(id uuid.UUID) error
	DayEvens(time time.Time) ([]*Event, error)
	WeekEvens(time time.Time) ([]*Event, error)
	MonthEvens(time time.Time) ([]*Event, error)
	GetEvensToNotify(limit int) ([]*EventNotification, error)
	Clean(olderThan time.Duration) error
}
