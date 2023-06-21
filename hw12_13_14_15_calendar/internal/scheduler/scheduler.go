package scheduler

import (
	"github.com/adrevin/ogdphw/hw12_13_14_15_calendar/internal/configuration"
	"github.com/adrevin/ogdphw/hw12_13_14_15_calendar/internal/logger"
	"github.com/adrevin/ogdphw/hw12_13_14_15_calendar/internal/mq"
	"github.com/adrevin/ogdphw/hw12_13_14_15_calendar/internal/storage"
)

const evensRequestLimit = 10

type Scheduler struct {
	logger  logger.Logger
	storage storage.Storage
	config  configuration.SchedulerConfiguration
	mq      mq.MQ
}

func New(
	logger logger.Logger,
	storage storage.Storage,
	config configuration.SchedulerConfiguration,
	mq mq.MQ,
) *Scheduler {
	app := &Scheduler{logger: logger, storage: storage, config: config, mq: mq}
	return app
}

func (s *Scheduler) Scan() error {
	events, err := s.storage.GetEvensToNotify(evensRequestLimit)
	if err != nil {
		return err
	}

	s.logger.Debugf("received %d events to notify", len(events))
	for _, event := range events {
		notification := &mq.Notification{ID: event.ID, Title: event.Title, Time: event.Time, UserID: event.UserID}

		err := s.mq.SendEventNotification(notification)
		if err != nil {
			s.logger.Errorf("Can not send notification: %+v", err)
			continue
		}
		s.logger.Debugf("Notification '%s' sent", notification.ID)

		err = s.storage.SetEvenIsNotified(notification.ID)
		if err != nil {
			s.logger.Errorf("Can not set event '%s' to is notified state: %+v", event.ID, err)
			continue
		}

		s.logger.Debugf("event '%s' notification sent", event.ID)
	}
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
