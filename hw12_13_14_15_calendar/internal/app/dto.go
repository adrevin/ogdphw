package app

import (
	"time"

	"github.com/google/uuid"
)

type EventRequest struct {
	Title    string    `json:"title" validate:"required"`
	Time     time.Time `json:"time" validate:"required"`
	Duration int64     `json:"duration" validate:"required"`
	UserID   uuid.UUID `json:"userId" validate:"required"`
}

type EventResponse struct {
	ID       uuid.UUID     `json:"id"`
	Title    string        `json:"title"`
	Time     time.Time     `json:"time"`
	Duration time.Duration `json:"duration"`
	UserID   uuid.UUID     `json:"userId"`
}

type EventID struct {
	ID uuid.UUID `json:"id"`
}
