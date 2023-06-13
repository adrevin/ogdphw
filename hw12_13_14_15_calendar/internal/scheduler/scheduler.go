package scheduler

import (
	"github.com/adrevin/ogdphw/hw12_13_14_15_calendar/internal/logger"
	"github.com/adrevin/ogdphw/hw12_13_14_15_calendar/internal/storage"
)

type Scheduler struct {
	logger  logger.Logger
	storage storage.Storage
}

func New(logger logger.Logger, storage storage.Storage) *Scheduler {
	app := &Scheduler{logger: logger, storage: storage}
	return app
}

func (s *Scheduler) Scan() error {
	return nil
}

func (s *Scheduler) Clean() error {
	return nil
}
