package app

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/adrevin/ogdphw/hw12_13_14_15_calendar/internal/logger"
	"github.com/adrevin/ogdphw/hw12_13_14_15_calendar/internal/storage"
	"github.com/go-http-utils/headers"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

const URLPattern = "/api/events/"

type App struct {
	logger   logger.Logger
	storage  storage.Storage
	validate *validator.Validate
}

func New(logger logger.Logger, storage storage.Storage) *App {
	v := validator.New()
	v.RegisterCustomTypeFunc(validateUUID, uuid.UUID{})
	app := &App{logger: logger, storage: storage, validate: v}
	logger.Info("calendar application created")
	return app
}

func validateUUID(field reflect.Value) interface{} {
	if valuer, ok := field.Interface().(uuid.UUID); ok {
		return valuer.String()
	}

	return nil
}

func (app *App) HandleCalendarRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		app.createEvent(w, r)
	case http.MethodGet:
		app.readEvents(w, r)
	case http.MethodPut:
		app.updateEvent(w, r)
	case http.MethodDelete:
		app.deleteEvent(w, r)
	default:
		http.Error(w, "not implemented", http.StatusNotImplemented)
	}
}

func (app *App) readEvents(w http.ResponseWriter, r *http.Request) {
	query := strings.Split(strings.ReplaceAll(r.RequestURI, URLPattern, ""), "/")
	if len(query) != 2 {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	command := query[0]
	if !contains([]string{"day", "week", "month"}, command) {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	t, err := time.Parse(time.DateOnly, query[1])

	// sanitize time
	t = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())

	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	switch command {
	case "day":
		evens, err := app.storage.DayEvens(t)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			app.logger.Errorf("can't get week evens: %+v", err)
			return
		}
		writeEvens(evens, app, w)
	case "week":
		evens, err := app.storage.WeekEvens(t)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			app.logger.Errorf("can't get week evens: %+v", err)
			return
		}
		writeEvens(evens, app, w)
	case "month":
		evens, err := app.storage.MonthEvens(t)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			app.logger.Errorf("can't get week evens: %+v", err)
			return
		}
		writeEvens(evens, app, w)
	default:
		http.Error(w, "bad request", http.StatusBadRequest)
	}
}

func writeEvens(evens []*storage.Event, app *App, w http.ResponseWriter) {
	response := make([]EventResponse, 0, len(evens))
	for _, event := range evens {
		// TODO: use mapping
		er := EventResponse{
			ID:       event.ID,
			Title:    event.Title,
			Time:     event.Time,
			Duration: event.Duration,
			UserID:   event.UserID,
		}
		response = append(response, er)
	}
	body, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		app.logger.Errorf("can't marshal evens: %+v", err)
		return
	}
	w.Header().Set(headers.ContentType, "application/json")
	// w.WriteHeader(http.StatusOK)
	_, err = w.Write(body)
	if err != nil {
		app.logger.Errorf("can't write response body: %+v", err)
	}
}

func (app *App) createEvent(w http.ResponseWriter, r *http.Request) {
	eventRequest := app.getEventRequest(w, r)
	if eventRequest == nil {
		return
	}
	e := &storage.Event{
		Title:    eventRequest.Title,
		Time:     eventRequest.Time,
		Duration: time.Duration(eventRequest.Duration) * time.Second,
		UserID:   eventRequest.UserID,
	}
	eventID, err := app.storage.Create(e)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		app.logger.Errorf("can't create event: %+v", err)
		return
	}

	body, err := json.Marshal(EventID{ID: eventID})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		app.logger.Errorf("can't marshal event response: %+v", err)
		return
	}
	w.Header().Set(headers.ContentType, "application/json")
	_, err = w.Write(body)
	if err != nil {
		app.logger.Errorf("can't write response: %+v", err)
	}
}

func (app *App) updateEvent(w http.ResponseWriter, r *http.Request) {
	eventID := app.getURLEventID(w, r)
	if eventID == nil {
		return
	}
	eventRequest := app.getEventRequest(w, r)
	if eventRequest == nil {
		return
	}
	storageEvent := &storage.Event{
		Title:    eventRequest.Title,
		Time:     eventRequest.Time,
		Duration: time.Duration(eventRequest.Duration) * time.Second,
		UserID:   eventRequest.UserID,
	}
	err := app.storage.Update(*eventID, storageEvent)
	if errors.Is(err, storage.ErrEventNotFound) {
		http.Error(w, "event not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (app *App) deleteEvent(w http.ResponseWriter, r *http.Request) {
	eventID := app.getURLEventID(w, r)
	if eventID == nil {
		return
	}
	err := app.storage.Delete(*eventID)
	if errors.Is(err, storage.ErrEventNotFound) {
		http.Error(w, "event not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func (app *App) getURLEventID(w http.ResponseWriter, r *http.Request) *uuid.UUID {
	idURLPart := strings.ReplaceAll(r.RequestURI, URLPattern, "")
	var eventID uuid.UUID
	eventID, err := uuid.Parse(idURLPart)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return nil
	}
	return &eventID
}

func (app *App) getEventRequest(w http.ResponseWriter, r *http.Request) *EventRequest {
	if r.Header.Get("Content-type") != "application/json" {
		http.Error(w, "bad request", http.StatusBadRequest)
		return nil
	}

	content, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		app.logger.Errorf("can't read request body: %+v", err)
		return nil
	}

	var request *EventRequest
	err = json.Unmarshal(content, &request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		app.logger.Errorf("can't unmarshal request body: %+v", err)
		return nil
	}

	err = app.validate.Struct(request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return nil
	}
	return request
}
