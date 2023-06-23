package sqlstorage

import (
	"database/sql"
	"os"
	"path"
	"sort"
	"time"

	"github.com/adrevin/ogdphw/hw12_13_14_15_calendar/internal/configuration"
	"github.com/adrevin/ogdphw/hw12_13_14_15_calendar/internal/logger"
	"github.com/adrevin/ogdphw/hw12_13_14_15_calendar/internal/storage"
	"github.com/google/uuid"
	// _ added to prevent error: unknown driver "postgres" (forgotten import?)
	_ "github.com/lib/pq"
	"github.com/snabb/isoweek"
)

type sqlStorage struct {
	config configuration.StorageConfiguration
	logger logger.Logger
	db     *sql.DB
}

func New(config configuration.StorageConfiguration, logger logger.Logger) storage.Storage {
	db, err := sql.Open("postgres", config.PostgresConnection)
	if err != nil {
		logger.Fatalf("can not open database connection: %+v", err)
	}
	err = db.Ping()
	if err != nil {
		logger.Fatalf("can not ping database: %+v", err)
	}

	logger.Debug("connected to database")
	return &sqlStorage{config: config, logger: logger, db: db}
}

func (s *sqlStorage) Create(event *storage.Event) (uuid.UUID, error) {
	dayKey := dayKey(event.Time)
	weekKey := weekKey(event.Time)
	monthKey := monthKey(event.Time)

	command := `insert into events (title, time, duration, user_id, day_key, week_key, month_key)
values ($1, $2, $3, $4, $5, $6, $7)
returning id`

	event.Time = event.Time.Truncate(time.Second)
	var id *uuid.UUID
	row := s.db.QueryRow(
		command,
		event.Title,
		event.Time,
		int64(event.Duration.Seconds()),
		event.UserID,
		dayKey,
		weekKey,
		monthKey)

	err := row.Scan(&id)
	if err != nil {
		s.logger.Errorf("can not insert event: %+v", err)
		return uuid.Nil, err
	}
	return *id, nil
}

func (s *sqlStorage) Update(id uuid.UUID, event *storage.Event) error {
	dayKey := dayKey(event.Time)
	weekKey := weekKey(event.Time)
	monthKey := monthKey(event.Time)
	command := `update events
set title=$1, time=$2, duration=$3, user_id=$4, day_key=$5, week_key=$6, month_key=$7
where id =$8`

	res, err := s.db.Exec(
		command, event.Title, event.Time, int64(event.Duration.Seconds()), event.UserID, dayKey, weekKey, monthKey, id)
	if err != nil {
		s.logger.Errorf("can not update event: %+v", err)
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		s.logger.Errorf("can not update event: %+v", err)
		return err
	}
	if rowsAffected != 1 {
		return storage.ErrEventNotFound
	}
	return nil
}

func (s *sqlStorage) Delete(id uuid.UUID) error {
	res, err := s.db.Exec("delete from events where id =$1", id)
	if err != nil {
		s.logger.Errorf("can not delete event: %+v", err)
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		s.logger.Errorf("can not delete event: %+v", err)
		return err
	}
	if rowsAffected != 1 {
		return storage.ErrEventNotFound
	}
	return nil
}

func (s *sqlStorage) DayEvens(t time.Time) ([]*storage.Event, error) {
	dayKey := dayKey(t)
	rows, err := s.db.Query("select id, title, time, duration, user_id from events where day_key=$1", dayKey)
	if err != nil {
		s.logger.Errorf("can not get events: %+v", err)
		return nil, err
	}
	events, err := getEvents(rows)
	if err != nil {
		s.logger.Errorf("can not get events: %+v", err)
		return nil, err
	}
	return events, nil
}

func (s *sqlStorage) WeekEvens(t time.Time) ([]*storage.Event, error) {
	weekKey := weekKey(t)
	rows, err := s.db.Query("select id, title, time, duration, user_id from events where week_key=$1", weekKey)
	if err != nil {
		s.logger.Errorf("can not get events: %+v", err)
		return nil, err
	}
	events, err := getEvents(rows)
	if err != nil {
		s.logger.Errorf("can not get events: %+v", err)
		return nil, err
	}
	return events, nil
}

func (s *sqlStorage) MonthEvens(t time.Time) ([]*storage.Event, error) {
	monthKey := monthKey(t)
	rows, err := s.db.Query("select id, title, time, duration, user_id from events where month_key=$1", monthKey)
	if err != nil {
		s.logger.Errorf("can not get events: %+v", err)
		return nil, err
	}
	events, err := getEvents(rows)
	if err != nil {
		s.logger.Errorf("can not get events: %+v", err)
		return nil, err
	}
	return events, nil
}

func (s *sqlStorage) GetEvensToNotify(limit int) ([]*storage.Event, error) {
	query := `select id, title, time, user_id from events where notified_at is null limit $1`

	rows, err := s.db.Query(query, limit)
	if err != nil {
		s.logger.Errorf("can not get events: %+v", err)
		return nil, err
	}

	events := make([]*storage.Event, 0)
	for rows.Next() {
		event := &storage.Event{}
		err := rows.Scan(&event.ID, &event.Title, &event.Time, &event.UserID)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}

func (s *sqlStorage) SetEvenIsNotified(eventID uuid.UUID) error {
	exec, err := s.db.Exec("update events set notified_at=$1 where id=$2", time.Now(), eventID)
	if err != nil {
		s.logger.Errorf("can not update event: %+v", err)
		return err
	}
	result, err := exec.RowsAffected()
	if err != nil {
		s.logger.Errorf("can not get rows affected : %+v", err)
		return err
	}
	if result == 0 {
		s.logger.Errorf("can not update row. row not found", err)
		return storage.ErrEventNotFound
	}
	return nil
}

func (s *sqlStorage) Clean(duration time.Duration) (int64, error) {
	max := time.Now().Add(-duration)
	r, err := s.db.Exec("delete from events where time < $1", max)
	if err != nil {
		return 0, err
	}
	affected, err := r.RowsAffected()
	if err != nil {
		return 0, err
	}
	return affected, nil
}

func getEvents(rows *sql.Rows) ([]*storage.Event, error) {
	events := make([]*storage.Event, 0)
	for rows.Next() {
		event := &storage.Event{}
		var d *int64
		err := rows.Scan(&event.ID, &event.Title, &event.Time, &d, &event.UserID)
		if err != nil {
			return nil, err
		}
		event.Duration = time.Duration(*d) * time.Second
		events = append(events, event)
	}
	return events, nil
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

func MigrateDatabase(config configuration.StorageConfiguration, logger logger.Logger) {
	db, err := sql.Open("postgres", config.PostgresConnection)
	if err != nil {
		logger.Fatalf("can not open database connection: %+v", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		logger.Fatalf("can not ping database: %+v", err)
		return
	}

	ex, err := os.Executable()
	if err != nil {
		logger.Fatalf("can not get executable: %+v", err)
		return
	}
	migrationsDir := path.Join(path.Dir(ex), "migrations")
	files, err := os.ReadDir(migrationsDir)
	if err != nil {
		logger.Fatalf("can not read dir %s: %+v", migrationsDir, err)
		return
	}

	sort.Slice(files, func(i, j int) bool {
		return files[i].Name() < files[j].Name()
	})

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		fileInfo, err := file.Info()
		_ = fileInfo
		if err != nil {
			logger.Fatalf("can not read file: %+v", err)
			return
		}
		fileName := file.Name()
		bytes, err := os.ReadFile(path.Join(migrationsDir, fileName))
		if err != nil {
			logger.Fatalf("can not read file: %+v", err)
			return
		}
		_, err = db.Exec(string(bytes))
		if err != nil {
			logger.Fatalf("can not exec '%s': %+v", fileName, err)
			return
		}
		logger.Debugf("applied migration '%s'", fileName)
	}
}
