package backgroundtask

import (
	"context"
	"time"

	"github.com/adrevin/ogdphw/hw12_13_14_15_calendar/internal/logger"
)

// TODO: add graceful stop.
type BackgroundTask struct {
	name    string
	delay   time.Duration
	context context.Context
	worker  func()
	logger  logger.Logger
}

func New(context context.Context, name string, delay time.Duration, worker func(), logger logger.Logger,
) *BackgroundTask {
	return &BackgroundTask{
		context: context,
		name:    name,
		delay:   delay,
		worker:  worker,
		logger:  logger,
	}
}

func (t *BackgroundTask) Start() {
	t.logger.Infof("process '%s' will run periodically with %s delay", t.name, t.delay)
	for {
		select {
		case <-t.context.Done():
			t.logger.Infof("process '%s' stopped", t.name)
			return
		case <-time.After(t.delay):
			t.worker()
		}
	}
}
