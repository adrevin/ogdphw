package mq

import (
	"time"

	"github.com/google/uuid"
)

type Notification struct {
	ID     uuid.UUID
	Title  string
	Time   time.Time
	UserID uuid.UUID
}
