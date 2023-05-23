package sqlstorage

import (
	"context"
	"time"

	"github.com/adrevin/ogdphw/hw12_13_14_15_calendar/internal/storage"
	"github.com/google/uuid"
)

type sqlStorage struct { //nolint:unused
	// TODO
}

func (s *sqlStorage) Create(event *storage.Event) uuid.UUID { //nolint:revive,unused
	// TODO implement me
	panic("implement me")
}

func (s *sqlStorage) Update(id uuid.UUID, event *storage.Event) error { //nolint:revive,unused
	// TODO implement me
	panic("implement me")
}

func (s *sqlStorage) Delete(id uuid.UUID) error { //nolint:revive,unused
	// TODO implement me
	panic("implement me")
}

func (s *sqlStorage) DayEvens(time time.Time) []*storage.Event { //nolint:revive,unused
	// TODO implement me
	panic("implement me")
}

func (s *sqlStorage) WeekEvens(time time.Time) []*storage.Event { //nolint:revive,unused
	// TODO implement me
	panic("implement me")
}

func (s *sqlStorage) MonthEvens(time time.Time) []*storage.Event { //nolint:revive,unused
	// TODO implement me
	panic("implement me")
}

func New() storage.Storage {
	// TODO implement me
	// return &sqlStorage{}
	panic("sql storage does not implemented")
}

func (s *sqlStorage) Connect(ctx context.Context) error { //nolint:revive,unused
	// TODO
	return nil
}

func (s *sqlStorage) Close(ctx context.Context) error { //nolint:revive,unused
	// TODO
	return nil
}
