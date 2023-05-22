package memorystorage

import (
	"errors"
	"sync"
	"time"

	"github.com/adrevin/ogdphw/hw12_13_14_15_calendar/internal/entities"
	"github.com/google/uuid"
	"github.com/snabb/isoweek"
)

type Storage struct {
	mu     *sync.RWMutex
	events map[uuid.UUID]*entities.Event
	days   map[time.Time]map[uuid.UUID]*entities.Event
	weeks  map[time.Time]map[uuid.UUID]*entities.Event
	months map[time.Time]map[uuid.UUID]*entities.Event
}

func New() *Storage {
	return &Storage{
		events: make(map[uuid.UUID]*entities.Event),
		days:   make(map[time.Time]map[uuid.UUID]*entities.Event),
		weeks:  make(map[time.Time]map[uuid.UUID]*entities.Event),
		months: make(map[time.Time]map[uuid.UUID]*entities.Event),
		mu:     &sync.RWMutex{},
	}
}

var ErrEventNotFound = errors.New("unsupported file")

func (l Storage) Create(event *entities.Event) uuid.UUID {
	defer l.mu.RUnlock()

	l.mu.RLock()
	event.ID = uuid.New()
	l.save(event)

	return event.ID
}

func (l Storage) Update(id uuid.UUID, event *entities.Event) error {
	defer l.mu.RUnlock()

	l.mu.RLock()
	if l.events[id] == nil {
		return ErrEventNotFound
	}
	event.ID = id
	l.save(event)
	return nil
}

func (l Storage) Delete(id uuid.UUID) error {
	defer l.mu.RUnlock()

	l.mu.RLock()
	event := l.events[id]
	if l.events[id] == nil {
		return ErrEventNotFound
	}

	delete(l.events, event.ID)

	dayKey := dayKey(event.Time)
	delete(l.days[dayKey], event.ID)

	monthKey := monthKey(event.Time)
	delete(l.months[monthKey], event.ID)

	weekKey := weekKey(event.Time)
	delete(l.weeks[weekKey], event.ID)

	return nil
}

func (l Storage) DayEvens(time time.Time) ([]entities.Event, error) { //nolint:revive
	// TODO
	return make([]entities.Event, 0), nil
}

func (l Storage) WeekEvens(time time.Time) ([]entities.Event, error) { //nolint:revive
	// TODO
	return make([]entities.Event, 0), nil
}

func (l Storage) MonthEvens(time time.Time) ([]entities.Event, error) { //nolint:revive
	// TODO
	return make([]entities.Event, 0), nil
}

func (l Storage) save(event *entities.Event) {
	l.events[event.ID] = event

	dayKey := dayKey(event.Time)
	if l.days[dayKey] == nil {
		l.days[dayKey] = make(map[uuid.UUID]*entities.Event)
	}
	l.days[dayKey][event.ID] = event

	weekKey := weekKey(event.Time)
	if l.weeks[weekKey] == nil {
		l.weeks[weekKey] = make(map[uuid.UUID]*entities.Event)
	}
	l.weeks[weekKey][event.ID] = event

	monthKey := monthKey(event.Time)
	if l.months[monthKey] == nil {
		l.months[monthKey] = make(map[uuid.UUID]*entities.Event)
	}
}

var location = time.UTC

func dayKey(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, location)
}

func monthKey(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), 0, 0, 0, 0, 0, location)
}

func weekKey(t time.Time) time.Time {
	year, week := isoweek.FromDate(t.Year(), t.Month(), t.Day())
	return isoweek.StartTime(year, week, location)
}
