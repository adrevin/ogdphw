package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/adrevin/ogdphw/hw12_13_14_15_calendar/internal/configuration"
	"github.com/adrevin/ogdphw/hw12_13_14_15_calendar/internal/logger"
	"github.com/adrevin/ogdphw/hw12_13_14_15_calendar/internal/mq"
	"github.com/adrevin/ogdphw/hw12_13_14_15_calendar/internal/mq/rabbitmq"
	sqlstorage "github.com/adrevin/ogdphw/hw12_13_14_15_calendar/internal/storage/sql"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "/etc/calendar/config.yml", "Path to configuration yaml file")
}

type notification struct {
	Title string
	Time  time.Time
}

func main() {
	flag.Parse()

	config, err := configuration.NewConfig(configFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	logger := logger.New(config.Logger)
	defer logger.Sync()

	rmq, err := rabbitmq.New(config.MessageQueue, logger)
	if err != nil {
		logger.Errorf("failed to start sender. RMQ error")
		os.Exit(1) //nolint:gocritic
	}
	defer rmq.Close()

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	logger.Infof("sender started")

	storageStorage := sqlstorage.New(config.Storage, logger)

	err = rmq.ConsumeNotifications(
		ctx,
		func(mqNotification *mq.Notification) bool {
			notification, err := json.Marshal(notification{Title: mqNotification.Title, Time: mqNotification.Time})
			if err != nil {
				logger.Errorf("failed to marshal notification: %+v", err)
				return false
			}

			err = storageStorage.RegisterNotification(
				mqNotification.UserID,
				mqNotification.ID,
				string(notification))

			if err != nil {
				logger.Errorf("failed to notify user '%s' about event '%s': %+v", mqNotification.UserID, mqNotification.ID, err)
				return false
			}

			logger.Debugf("user '%s' notified about event '%s'", mqNotification.UserID, mqNotification.ID)
			return true
		})
	if err != nil {
		logger.Errorf("consuming error", "%+v", err)
	}

	logger.Infof("sender stopped")
}
