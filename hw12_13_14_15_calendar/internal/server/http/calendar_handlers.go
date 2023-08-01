package internalhttp

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/adrevin/ogdphw/hw12_13_14_15_calendar/internal/storage"
	"github.com/go-http-utils/headers"
	"github.com/google/uuid"
)

const URLPattern = "/api/events/"

func (s *Server) handleCalendarRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		s.createEvent(w, r)
	case http.MethodGet:
		s.readEvents(w, r)
	case http.MethodPut:
		s.updateEvent(w, r)
	case http.MethodDelete:
		s.deleteEvent(w, r)
	default:
		http.Error(w, "not implemented", http.StatusNotImplemented)
	}
}

func (s *Server) createEvent(w http.ResponseWriter, r *http.Request) {
	eventRequest := s.getEventRequest(w, r)
	if eventRequest == nil {
		return
	}
	e := &storage.Event{
		Title:    eventRequest.Title,
		Time:     eventRequest.Time,
		Duration: time.Duration(eventRequest.Duration) * time.Second,
		UserID:   eventRequest.UserID,
	}
	eventID, err := s.app.CreateEvent(e)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		s.logger.Errorf("can't create event: %+v", err)
		return
	}

	body, err := json.Marshal(EventID{ID: eventID})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		s.logger.Errorf("can't marshal event response: %+v", err)
		return
	}
	w.Header().Set(headers.ContentType, "application/json")
	_, err = w.Write(body)
	if err != nil {
		s.logger.Errorf("can't write response: %+v", err)
	}
}

func (s *Server) readEvents(w http.ResponseWriter, r *http.Request) {
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
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	// sanitize time
	t = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())

	switch command {
	case "day":
		evens, err := s.app.DayEvens(t)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			s.logger.Errorf("can't get week evens: %+v", err)
			return
		}
		s.writeEvens(evens, w)
	case "week":
		evens, err := s.app.WeekEvens(t)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			s.logger.Errorf("can't get week evens: %+v", err)
			return
		}
		s.writeEvens(evens, w)
	case "month":
		evens, err := s.app.MonthEvens(t)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			s.logger.Errorf("can't get week evens: %+v", err)
			return
		}
		s.writeEvens(evens, w)
	default:
		http.Error(w, "bad request", http.StatusBadRequest)
	}
}

func (s *Server) updateEvent(w http.ResponseWriter, r *http.Request) {
	eventID := s.getURLEventID(w, r)
	if eventID == nil {
		return
	}
	eventRequest := s.getEventRequest(w, r)
	if eventRequest == nil {
		return
	}
	storageEvent := &storage.Event{
		Title:    eventRequest.Title,
		Time:     eventRequest.Time,
		Duration: time.Duration(eventRequest.Duration) * time.Second,
		UserID:   eventRequest.UserID,
	}
	err := s.app.UpdateEvent(*eventID, storageEvent)
	if errors.Is(err, storage.ErrEventNotFound) {
		http.Error(w, "event not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Server) deleteEvent(w http.ResponseWriter, r *http.Request) {
	eventID := s.getURLEventID(w, r)
	if eventID == nil {
		return
	}
	err := s.app.DeleteEvent(*eventID)
	if errors.Is(err, storage.ErrEventNotFound) {
		http.Error(w, "event not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Server) getURLEventID(w http.ResponseWriter, r *http.Request) *uuid.UUID {
	idURLPart := strings.ReplaceAll(r.RequestURI, URLPattern, "")
	var eventID uuid.UUID
	eventID, err := uuid.Parse(idURLPart)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return nil
	}
	return &eventID
}

func (s *Server) getEventRequest(w http.ResponseWriter, r *http.Request) *EventRequest {
	if r.Header.Get("Content-type") != "application/json" {
		http.Error(w, "bad request", http.StatusBadRequest)
		return nil
	}

	content, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		s.logger.Errorf("can't read request body: %+v", err)
		return nil
	}

	var request *EventRequest
	err = json.Unmarshal(content, &request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		s.logger.Errorf("can't unmarshal request body: %+v", err)
		return nil
	}

	err = s.validate.Struct(request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return nil
	}
	return request
}

func (s *Server) writeEvens(evens []*storage.Event, w http.ResponseWriter) {
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
		s.logger.Errorf("can't marshal evens: %+v", err)
		return
	}
	w.Header().Set(headers.ContentType, "application/json")
	_, err = w.Write(body)
	if err != nil {
		s.logger.Errorf("can't write response body: %+v", err)
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
