package app

import (
	"time"

	"github.com/adrevin/ogdphw/hw12_13_14_15_calendar/internal/logger"
	"github.com/adrevin/ogdphw/hw12_13_14_15_calendar/internal/storage"
	"github.com/google/uuid"
)

const URLPattern = "/api/events/"

type App struct {
	logger  logger.Logger
	storage storage.Storage
}

func New(logger logger.Logger, storage storage.Storage) App {
	app := App{logger: logger, storage: storage}
	logger.Info("calendar application created")
	return app
}

func (app *App) CreateEvent(event *storage.Event) (uuid.UUID, error) {
	eventID, err := app.storage.Create(event)
	if err != nil {
		app.logger.Errorf("can't create event: %v+", err)
	} else {
		app.logger.Debugf("created event: %s", eventID)
	}
	return eventID, err
}

func (app *App) UpdateEvent(eventID uuid.UUID, event *storage.Event) error {
	err := app.storage.Update(eventID, event)
	if err != nil {
		app.logger.Errorf("can't update event '%s': %v+", eventID, err)
	} else {
		app.logger.Debugf("updated event: %s", eventID)
	}
	return err
}

func (app *App) DeleteEvent(eventID uuid.UUID) error {
	err := app.storage.Delete(eventID)
	if err != nil {
		app.logger.Errorf("can't delete event '%s': %v+", eventID, err)
		return err
	}
	app.logger.Debugf("deleted event: %s", eventID)
	return err
}

func (app *App) DayEvens(t time.Time) ([]*storage.Event, error) {
	events, err := app.storage.DayEvens(t)
	if err != nil {
		app.logger.Errorf("can't get day events: %v+", err)
	}
	return events, err
}

func (app *App) WeekEvens(t time.Time) ([]*storage.Event, error) {
	events, err := app.storage.WeekEvens(t)
	if err != nil {
		app.logger.Errorf("can't get week events: %v+", err)
	}
	return events, err
}

func (app *App) MonthEvens(t time.Time) ([]*storage.Event, error) {
	events, err := app.storage.MonthEvens(t)
	if err != nil {
		app.logger.Errorf("can't get week events: %v+", err)
	}
	return events, err
}
