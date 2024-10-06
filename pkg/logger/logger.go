package logger

import (
	"go.uber.org/zap"
)

var logger *zap.SugaredLogger

func Get() *zap.SugaredLogger {
	return logger
}

func Initialize() (*zap.SugaredLogger, error) {
	baseLogger, err := zap.NewProduction()

	if err != nil {
		return nil, err
	}

	logger = baseLogger.Sugar()

	defer baseLogger.Sync()

	return logger, nil
}
