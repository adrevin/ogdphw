package storage

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	ID       uuid.UUID
	Title    string
	Time     time.Time
	Duration time.Duration
	UserID   uuid.UUID
	DayKey   time.Time
	WeekKey  time.Time
	MonthKey time.Time
}

type EventNotification struct {
	ID     uuid.UUID
	Title  string
	Time   time.Time
	UserID uuid.UUID
}
