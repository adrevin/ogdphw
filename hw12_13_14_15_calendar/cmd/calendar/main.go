package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/adrevin/ogdphw/hw12_13_14_15_calendar/internal/app"
	"github.com/adrevin/ogdphw/hw12_13_14_15_calendar/internal/configuration"
	"github.com/adrevin/ogdphw/hw12_13_14_15_calendar/internal/logger"
	internalhttp "github.com/adrevin/ogdphw/hw12_13_14_15_calendar/internal/server/http"
	"github.com/adrevin/ogdphw/hw12_13_14_15_calendar/internal/storage"
	memorystorage "github.com/adrevin/ogdphw/hw12_13_14_15_calendar/internal/storage/memory"
	sqlstorage "github.com/adrevin/ogdphw/hw12_13_14_15_calendar/internal/storage/sql"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "/etc/calendar/config.yml", "Path to configuration yaml file")
}

func main() {
	flag.Parse()

	if flag.Arg(0) == "version" {
		printVersion()
		return
	}

	config, err := configuration.NewConfig(configFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	logg := logger.New(config.Logger)
	defer logg.Sync()

	var storageStorage storage.Storage
	if config.Storage.UsePostgresStorage {
		storageStorage = sqlstorage.New()
	} else {
		storageStorage = memorystorage.New()
	}
	calendar := app.New(logg, storageStorage)

	server := internalhttp.NewServer(logg, calendar, config.Server)

	ctx, cancel := signal.NotifyContext(context.Background(),
		syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	go func() {
		<-ctx.Done()

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		if err := server.Stop(ctx); err != nil {
			logg.Error("failed to stop http server: " + err.Error())
		}
	}()

	logg.Info("calendar is running...")

	if err := server.Start(ctx); err != nil {
		logg.Error("failed to start http server: " + err.Error())
		cancel()
		os.Exit(1) //nolint:gocritic
	}
}
