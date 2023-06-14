package scheduler

import (
	"github.com/adrevin/ogdphw/hw12_13_14_15_calendar/internal/configuration"
	"github.com/adrevin/ogdphw/hw12_13_14_15_calendar/internal/logger"
	"github.com/adrevin/ogdphw/hw12_13_14_15_calendar/internal/storage"
)

type Scheduler struct {
	logger  logger.Logger
	storage storage.Storage
	config  configuration.SchedulerConfiguration
}

func New(logger logger.Logger, storage storage.Storage, config configuration.SchedulerConfiguration) *Scheduler {
	app := &Scheduler{logger: logger, storage: storage, config: config}
	return app
}

func (s *Scheduler) Scan() error {
	return nil
}

func (s *Scheduler) Clean() error {
	count, err := s.storage.Clean(s.config.CleanOlderThan)
	if err != nil {
		s.logger.Errorf("Can not delete old events: %+v", err)
		return err
	}
	s.logger.Debugf("Deleted %d old events", count)
	return nil
}
