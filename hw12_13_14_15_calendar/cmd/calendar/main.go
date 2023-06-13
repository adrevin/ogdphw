package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/adrevin/ogdphw/hw12_13_14_15_calendar/internal/app"
	"github.com/adrevin/ogdphw/hw12_13_14_15_calendar/internal/configuration"
	"github.com/adrevin/ogdphw/hw12_13_14_15_calendar/internal/logger"
	internalgrpc "github.com/adrevin/ogdphw/hw12_13_14_15_calendar/internal/server/grpc"
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

	if flag.Arg(0) == "migrate-db" {
		logg.Info("starting database migration ...")
		sqlstorage.MigrateDatabase(config.Storage, logg)
		logg.Info("database migration done")
		return
	}

	var storageStorage storage.Storage
	if config.Storage.UsePostgresStorage {
		storageStorage = sqlstorage.New(config.Storage, logg)
	} else {
		storageStorage = memorystorage.New()
	}

	calendar := app.New(logg, storageStorage)

	httpServer := internalhttp.NewServer(logg, calendar, config.HTTPServer)
	grpcServer := internalgrpc.NewServer(logg, calendar, config.GrpcServer)

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	go func() {
		<-ctx.Done()
		ctx, cancel := context.WithTimeout(context.Background(), config.ShutdownTimeout)
		defer cancel()
		go func() {
			if err := httpServer.Stop(ctx); err != nil {
				logg.Errorf("failed to stop http server: %+v", err)
			}
		}()
		go func() {
			grpcServer.Stop()
		}()
	}()

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		if err := httpServer.Start(ctx); err != nil {
			logg.Error("failed to start http server: " + err.Error())
			cancel()
			os.Exit(1)
		}
	}()

	go func() {
		defer wg.Done()
		if err := grpcServer.Start(ctx); err != nil {
			logg.Error("failed to start grpc server: " + err.Error())
			cancel()
			os.Exit(1)
		}
	}()

	logg.Info("calendar is running")
	wg.Wait()
}
