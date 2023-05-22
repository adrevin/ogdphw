package entities

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

type Notification struct {
	EventID     uuid.UUID
	EventTitle  string
	EventTime   time.Time
	EventUserID uuid.UUID
}
