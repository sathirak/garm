package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/sathirak/garm/internal/config"
	"github.com/sathirak/garm/pkg/logger"
)

var db *sql.DB

func Get() *sql.DB {
	return db
}

func Close() {
	log := logger.Get()

	err := db.Close()

	if err != nil {
		log.Errorw("shutdown", "package", "db", "error", err.Error())
		return
	}
	log.Infow("shutdown", "package", "db", "status", "ok")
}

func Initialize() {

	log := logger.Get()

	cfg := config.Get()

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		cfg.Database.Host, cfg.Database.Port, cfg.Database.Username, cfg.Database.Password, cfg.Database.DBName)

	var err error

	db, err = sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Errorw("startup", "package", "db", "error", err.Error())
	}

	err = db.Ping()

	if err != nil {
		log.Errorw("startup", "package", "db", "error", err.Error())
	}

	log.Infow("startup", "package", "db", "status", "ok")
}
