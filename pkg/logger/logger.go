package logger

import (
	"log"

	"go.uber.org/zap"
)

var logger *zap.SugaredLogger
var base *zap.Logger

func Get() *zap.SugaredLogger {
	return logger
}

func Initialize() {
	base, err := zap.NewProduction()

	if err != nil {
		log.Printf(`{"startup": "package", "logger": "error", "message": "%s"}`, err.Error())
		return
	}

	logger = base.Sugar()

	log.Print(`{"startup": "package", "logger": "status", "message": "ok"}`)
}

func Close() {
	err := logger.Sync()
	if err != nil {
		log.Printf(`{"startup": "package", "logger": "error", "message": "%s"}`, err.Error())
		return
	}
	err = base.Sync()
	if err != nil {
		log.Printf(`{"startup": "package", "logger": "error", "message": "%s"}`, err.Error())
		return
	}
}
