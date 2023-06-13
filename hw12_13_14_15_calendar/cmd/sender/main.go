package main

import (
	"flag"
	"fmt"
	"os"

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
	// TODO
}
