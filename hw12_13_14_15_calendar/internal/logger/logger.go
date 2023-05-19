package logger

import (
	"fmt"

	"github.com/adrevin/ogdphw/hw12_13_14_15_calendar/internal/configuration"
)

var cfg configuration.LoggerConf

type Logger struct { // TODO
}

func New(loggerConf configuration.LoggerConf) *Logger {
	cfg = loggerConf
	return &Logger{}
}

func (l Logger) Info(msg string) {
	fmt.Println(msg)
}

func (l Logger) Error(msg string) {
	// TODO
}

// TODO
