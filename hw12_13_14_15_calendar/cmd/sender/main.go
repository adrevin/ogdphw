package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/adrevin/ogdphw/hw12_13_14_15_calendar/internal/configuration"
	"github.com/adrevin/ogdphw/hw12_13_14_15_calendar/internal/logger"
	"github.com/adrevin/ogdphw/hw12_13_14_15_calendar/internal/mq"
	"github.com/adrevin/ogdphw/hw12_13_14_15_calendar/internal/mq/rabbitmq"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "/etc/calendar/config.yml", "Path to configuration yaml file")
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

	rmq := rabbitmq.New(config.MessageQueue, logger)
	defer rmq.Close()

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	logger.Debugf("sender started")
	err = rmq.ConsumeNotifications(
		ctx,
		func(notification *mq.Notification) bool {
			logger.Debugf("user '%s' notified about event '%s'", notification.UserID, notification.ID)
			return true
		})
	if err != nil {
		logger.Errorf("consuming error", "%+v", err)
	}

	logger.Debugf("sender stopped")
}
