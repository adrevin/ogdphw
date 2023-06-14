package memorystorage

import (
	"sync"
	"time"

	"github.com/adrevin/ogdphw/hw12_13_14_15_calendar/internal/storage"
	"github.com/google/uuid"
	"github.com/snabb/isoweek"
)

type memStorage struct {
	mu     *sync.RWMutex
	events map[uuid.UUID]*storage.Event
	days   map[time.Time]map[uuid.UUID]*storage.Event
	weeks  map[time.Time]map[uuid.UUID]*storage.Event
	months map[time.Time]map[uuid.UUID]*storage.Event
}

func New() storage.Storage {
	return &memStorage{
		events: make(map[uuid.UUID]*storage.Event),
		days:   make(map[time.Time]map[uuid.UUID]*storage.Event),
		weeks:  make(map[time.Time]map[uuid.UUID]*storage.Event),
		months: make(map[time.Time]map[uuid.UUID]*storage.Event),
		mu:     &sync.RWMutex{},
	}
}

func (l memStorage) Create(event *storage.Event) (uuid.UUID, error) {
	defer l.mu.RUnlock()

	l.mu.RLock()
	event.ID = uuid.New()
	setKeys(event)
	l.makeMaps(event)
	l.save(event)

	return event.ID, nil
}

func (l memStorage) Update(id uuid.UUID, event *storage.Event) error {
	defer l.mu.RUnlock()

	l.mu.RLock()

	if l.events[id] == nil {
		return storage.ErrEventNotFound
	}
	l.Delete(id)

	event.ID = id
	setKeys(event)
	l.makeMaps(event)
	l.save(event)

	return nil
}

func (l memStorage) Delete(id uuid.UUID) error {
	defer l.mu.RUnlock()

	l.mu.RLock()
	event := l.events[id]
	if l.events[id] == nil {
		return storage.ErrEventNotFound
	}

	delete(l.events, event.ID)

	delete(l.days[event.DayKey], event.ID)
	if len(l.days[event.DayKey]) == 0 {
		delete(l.days, event.DayKey)
	}

	delete(l.weeks[event.WeekKey], event.ID)
	if len(l.weeks[event.WeekKey]) == 0 {
		delete(l.weeks, event.WeekKey)
	}

	delete(l.months[event.MonthKey], event.ID)
	if len(l.months[event.MonthKey]) == 0 {
		delete(l.months, event.MonthKey)
	}

	return nil
}

func (l memStorage) DayEvens(time time.Time) ([]*storage.Event, error) {
	defer l.mu.RUnlock()

	l.mu.RLock()
	dayKey := dayKey(time)

	return eventsToResult(l.days[dayKey]), nil
}

func (l memStorage) WeekEvens(time time.Time) ([]*storage.Event, error) {
	defer l.mu.RUnlock()

	l.mu.RLock()
	weekKey := weekKey(time)

	return eventsToResult(l.weeks[weekKey]), nil
}

func (l memStorage) MonthEvens(time time.Time) ([]*storage.Event, error) {
	defer l.mu.RUnlock()

	l.mu.RLock()
	monthKey := monthKey(time)

	return eventsToResult(l.months[monthKey]), nil
}

func (memStorage) GetEvensToNotify(_ int) ([]*storage.EventNotification, error) {
	return nil, storage.ErrNotImplemented
}

func (memStorage) Clean(_ time.Duration) (int64, error) {
	return 0, storage.ErrNotImplemented
}

func eventsToResult(events map[uuid.UUID]*storage.Event) []*storage.Event {
	if events == nil {
		return make([]*storage.Event, 0)
	}
	result := make([]*storage.Event, 0, len(events))
	for _, event := range events {
		result = append(result, event)
	}
	return result
}

func (l memStorage) save(event *storage.Event) {
	l.events[event.ID] = event
	l.days[event.DayKey][event.ID] = event
	l.weeks[event.WeekKey][event.ID] = event
	l.months[event.MonthKey][event.ID] = event
}

func setKeys(event *storage.Event) {
	event.DayKey = dayKey(event.Time)
	event.WeekKey = weekKey(event.Time)
	event.MonthKey = monthKey(event.Time)
}

func (l memStorage) makeMaps(event *storage.Event) {
	if l.days[event.DayKey] == nil {
		l.days[event.DayKey] = make(map[uuid.UUID]*storage.Event)
	}

	if l.weeks[event.WeekKey] == nil {
		l.weeks[event.WeekKey] = make(map[uuid.UUID]*storage.Event)
	}

	if l.months[event.MonthKey] == nil {
		l.months[event.MonthKey] = make(map[uuid.UUID]*storage.Event)
	}
}

func dayKey(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

func monthKey(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), 0, 0, 0, 0, 0, t.Location())
}

func weekKey(t time.Time) time.Time {
	year, week := isoweek.FromDate(t.Year(), t.Month(), t.Day())
	return isoweek.StartTime(year, week, t.Location())
}
