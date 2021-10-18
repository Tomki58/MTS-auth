package common

import (
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

func init() {
	Logger = new(zap.Logger)
	var err error
	Logger, err = New()
	if err != nil {
		log.Fatal(err.Error())
	}
}

// New creates and return the new logger
func New() (*zap.Logger, error) {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "timestamp",
		LevelKey:       "level",
		NameKey:        "logger",
		MessageKey:     "message",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.RFC3339TimeEncoder,
		EncodeDuration: zapcore.MillisDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		CallerKey:      "caller",
		StacktraceKey:  "stacktrace",
	}

	loggerConfig := zap.Config{
		EncoderConfig:    encoderConfig,
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
		Encoding:         "json",
		Level:            zap.NewAtomicLevelAt(zap.InfoLevel),
	}

	logger, err := loggerConfig.Build()
	if err != nil {
		log.Fatal(err)
	}
	defer logger.Sync()
	return logger, nil
}
