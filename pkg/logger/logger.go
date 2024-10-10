package logger

import (
	"fmt"

	"go.uber.org/zap"
)

var logger *zap.SugaredLogger

func Get() *zap.SugaredLogger {
	return logger
}

func Initialize() {
	base, err := zap.NewProduction()

	if err != nil {
		fmt.Print("startup", "package", "logger", "error", err.Error())
		return
	}

	logger = base.Sugar()
	logger.Infow("startup", "package", "logger", "status", "ok")

	defer func() {
		if err := base.Sync(); err != nil {
			logger.Errorw("shutdown", "package", "logger", "error", err.Error())
		}
	}()
}
