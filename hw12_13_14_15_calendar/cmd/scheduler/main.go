package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/adrevin/ogdphw/hw12_13_14_15_calendar/internal/backgroundtask"
	"github.com/adrevin/ogdphw/hw12_13_14_15_calendar/internal/configuration"
	"github.com/adrevin/ogdphw/hw12_13_14_15_calendar/internal/logger"
	"github.com/adrevin/ogdphw/hw12_13_14_15_calendar/internal/scheduler"
	sqlstorage "github.com/adrevin/ogdphw/hw12_13_14_15_calendar/internal/storage/sql"
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

	if !config.Storage.UsePostgresStorage {
		fmt.Println("memory storage not implements required methods")
		os.Exit(1) //nolint:gocritic
	}

	storageStorage := sqlstorage.New(config.Storage, logger)
	scheduler := scheduler.New(logger, storageStorage)

	logger.Infof(
		"Scheduler started. Scan delay: %s, clean delay: %s",
		config.Scheduler.ScanDelay,
		config.Scheduler.CleanDelay)

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		backgroundtask.New(ctx, "Clean", config.Scheduler.CleanDelay, func() {
			err := scheduler.Clean()
			if err != nil {
				logger.Debugf("cleaning error: %=v", err)
				return
			}
		}, logger).Start()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		backgroundtask.New(ctx, "Scan", config.Scheduler.ScanDelay, func() {
			err := scheduler.Scan()
			if err != nil {
				logger.Debugf("cleaning error: %=v", err)
				return
			}
		}, logger).Start()
	}()

	wg.Wait()
	logger.Info("Scheduler stopped")
}
