package logger

import (
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger interface {
	Sync()
	Debug(msg string)
	Info(msg string)
	Warn(msg string)
	Error(msg string)
	Panic(msg string)
	Fatal(msg string)
}

type logg struct {
	logger *zap.Logger
}

func New(config zap.Config) Logger {
	// defaults
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig.EncodeLevel = zapcore.LowercaseLevelEncoder
	config.OutputPaths = []string{"stdout"}
	config.ErrorOutputPaths = []string{"stdout"}

	logger, err := config.Build()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return &logg{logger: logger}
}

func (l logg) Sync() {
	l.logger.Sync()
}

func (l logg) Debug(msg string) {
	l.logger.Debug(msg)
}

func (l logg) Info(msg string) {
	l.logger.Info(msg)
}

func (l logg) Warn(msg string) {
	l.logger.Warn(msg)
}

func (l logg) Error(msg string) {
	l.logger.Error(msg)
}

func (l logg) Panic(msg string) {
	l.logger.Panic(msg)
}

func (l logg) Fatal(msg string) {
	l.logger.Fatal(msg)
}
