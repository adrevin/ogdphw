package app

import (
	"context"
	"encoding/json"
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

func (app *App) CreateEvent(ctx context.Context, id, title string) error { //nolint:revive
	app.storage.Create(&storage.Event{Title: title})
	return nil
}

const EventsURLPattern = "/api/events/"

func (app *App) HandleCalendarRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	// create
	case http.MethodPost:
		if r.Header.Get("Content-type") != "application/json" {
			http.Error(w, "bad request", http.StatusBadRequest)
		}

		content, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		var request *EventRequest
		err = json.Unmarshal(content, &request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = app.validate.Struct(request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		e := &storage.Event{
			Title:    request.Title,
			Time:     request.Time,
			Duration: time.Duration(request.Duration) * time.Second,
			UserID:   request.UserID,
		}
		_, err = app.storage.Create(e)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	// read
	case http.MethodGet:
		app.handleReadRequest(w, r)
	// update
	case http.MethodPut:
	// update
	case http.MethodDelete:
	default:
		//		NotImplemented(w, r)
	}
}

func (app *App) handleReadRequest(w http.ResponseWriter, r *http.Request) {
	query := strings.Split(strings.ReplaceAll(r.RequestURI, EventsURLPattern, ""), "/")
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

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
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
