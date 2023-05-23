package entities

import (
	"time"

	"github.com/google/uuid"
)

type Notification struct {
	EventID     uuid.UUID
	EventTitle  string
	EventTime   time.Time
	EventUserID uuid.UUID
}
