package main

import (
	"github.com/sathirak/garm/internal/config"
	"fmt"
	"net/http"

	// "github.com/sathirak/garm/internal/db"
	"github.com/sathirak/garm/controllers"
	"github.com/sathirak/garm/pkg/logger"
)

func main() {
	log, err := logger.Initialize()

	if err != nil {
		log.Fatalf("error initializing logger")
	}

	cfg, err := config.Initialize()

	if err != nil {
		log.Fatalf("error initializing environment config")
	}

	// db.Initialize()

	mux := http.NewServeMux()

	mux.HandleFunc("GET /healthz", controllers.Healthz)

	err = http.ListenAndServe(fmt.Sprintf(":%s", cfg.App.Port), mux)

	if err != nil {
		log.Error(err.Error())
	}
}
