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
	Debugf(tmpl string, args ...interface{})
	Info(msg string)
	Infof(tmpl string, args ...interface{})
	Warn(msg string)
	Warnf(tmpl string, args ...interface{})
	Error(msg string)
	Errorf(tmpl string, args ...interface{})
	Panic(msg string)
	Panicf(tmpl string, args ...interface{})
	Fatal(msg string)
	Fatalf(tmpl string, args ...interface{})
}

type logg struct {
	logger *zap.Logger
	sugar  *zap.SugaredLogger
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
	return &logg{logger: logger, sugar: logger.Sugar()}
}

func (l logg) Sync() {
	l.logger.Sync()
	l.sugar.Sync()
}

func (l logg) Debug(msg string) {
	l.logger.Debug(msg)
}

func (l logg) Debugf(tmpl string, args ...interface{}) {
	l.sugar.Debugf(tmpl, args...)
}

func (l logg) Info(msg string) {
	l.logger.Info(msg)
}

func (l logg) Infof(tmpl string, args ...interface{}) {
	l.sugar.Infof(tmpl, args...)
}

func (l logg) Warn(msg string) {
	l.logger.Warn(msg)
}

func (l logg) Warnf(tmpl string, args ...interface{}) {
	l.sugar.Warnf(tmpl, args...)
}

func (l logg) Error(msg string) {
	l.logger.Error(msg)
}

func (l logg) Errorf(tmpl string, args ...interface{}) {
	l.sugar.Errorf(tmpl, args...)
}

func (l logg) Panic(msg string) {
	l.logger.Panic(msg)
}

func (l logg) Panicf(tmpl string, args ...interface{}) {
	l.sugar.Panicf(tmpl, args...)
}

func (l logg) Fatal(msg string) {
	l.logger.Fatal(msg)
}

func (l logg) Fatalf(tmpl string, args ...interface{}) {
	l.sugar.Fatalf(tmpl, args...)
}
