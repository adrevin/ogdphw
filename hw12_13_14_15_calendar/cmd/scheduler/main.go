package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	backgroundTask "github.com/adrevin/ogdphw/hw12_13_14_15_calendar/internal/backgroundtask"
	"github.com/adrevin/ogdphw/hw12_13_14_15_calendar/internal/configuration"
	"github.com/adrevin/ogdphw/hw12_13_14_15_calendar/internal/logger"
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
		backgroundTask.New(ctx, "Clean", config.Scheduler.CleanDelay, func() {
			logger.Debug("cleaning ...")
			time.Sleep(3 * time.Second)
			logger.Debug("cleaned")
		}, logger).Start()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		backgroundTask.New(ctx, "Scan", config.Scheduler.ScanDelay, func() {
			logger.Debug("scanning ...")
			time.Sleep(3 * time.Second)
			logger.Debug("scanned")
		}, logger).Start()
	}()

	wg.Wait()
}
