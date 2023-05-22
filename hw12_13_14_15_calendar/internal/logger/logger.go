package logger

import (
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	zapLogger *zap.Logger
}

func New(config zap.Config) *Logger {
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
	return &Logger{zapLogger: logger}
}

func (l Logger) Sync() {
	l.zapLogger.Sync()
}

func (l Logger) Debug(msg string, fields ...zap.Field) {
	l.zapLogger.Debug(msg, fields...)
}

func (l Logger) Info(msg string, fields ...zap.Field) {
	l.zapLogger.Info(msg, fields...)
}

func (l Logger) Warn(msg string, fields ...zap.Field) {
	l.zapLogger.Warn(msg, fields...)
}

func (l Logger) Error(msg string, fields ...zap.Field) {
	l.zapLogger.Error(msg, fields...)
}

func (l Logger) DPanic(msg string, fields ...zap.Field) {
	l.zapLogger.DPanic(msg, fields...)
}

func (l Logger) Panic(msg string, fields ...zap.Field) {
	l.zapLogger.Panic(msg, fields...)
}

func (l Logger) Fatal(msg string, fields ...zap.Field) {
	l.zapLogger.Fatal(msg, fields...)
}
