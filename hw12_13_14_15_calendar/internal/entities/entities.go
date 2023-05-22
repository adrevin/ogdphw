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
}

type Notification struct {
	EventID     uuid.UUID
	EventTitle  string
	EventTime   time.Time
	EventUserID uuid.UUID
}
