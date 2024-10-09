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
		fmt.Printf("error initializing logger:\n %v", err)
		return
	}

	logger = base.Sugar()
	logger.Infow("startup", "package", "logger", "status", "ok")

	defer base.Sync()
}
