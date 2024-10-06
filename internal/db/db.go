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

func Initialize() {

	log := logger.Get()

	cfg := config.Get()

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		cfg.Database.Host, cfg.Database.Port, cfg.Database.Username, cfg.Database.Password, cfg.Database.DBName)

	var err error

	db, err = sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Panic("error pinging db", err)
	}

	err = db.Ping()

	if err != nil {
		log.Panic("error connecting to db", err)
	}

	log.Info("db connection successful")
}
