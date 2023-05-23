package app

import (
	"context"

	"github.com/adrevin/ogdphw/hw12_13_14_15_calendar/internal/logger"
)

type App struct { // TODO
}

type Storage interface { // TODO
}

func New(logger logger.Logger, storage Storage) *App { //nolint:revive
	return &App{}
}

func (a *App) CreateEvent(ctx context.Context, id, title string) error { //nolint:revive
	// TODO
	return nil
	// return a.storage.CreateEvent(storage.Event{ID: id, Title: title})
}

// TODO
