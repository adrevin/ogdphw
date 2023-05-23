package storage

import (
	"time"

	"github.com/google/uuid"
)

type Storage interface {
	Create(event *Event) uuid.UUID
	Update(id uuid.UUID, event *Event) error
	Delete(id uuid.UUID) error
	DayEvens(time time.Time) []*Event
	WeekEvens(time time.Time) []*Event
	MonthEvens(time time.Time) []*Event
}
