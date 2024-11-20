package db

import (
	"database/sql"
	"fmt"

	"github.com/hotelbear/garm/internal/config"
	"github.com/hotelbear/garm/internal/logger"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *sql.DB
var gormDB *gorm.DB

func Get() *gorm.DB {
	return gormDB.Session(&gorm.Session{PrepareStmt: false})
}

func Close() {
	log := logger.Get()
	gormDB.Exec("DEALLOCATE ALL")
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

	gormDB, err = gorm.Open(postgres.Open(psqlInfo), &gorm.Config{PrepareStmt: true})

	if err != nil {
		log.Errorw("startup", "package", "db", "error", err.Error())
	}

	db, err = gormDB.DB()

	gormDB.Exec("DEALLOCATE ALL")

	if err != nil {
		log.Errorw("startup", "package", "db", "error", err.Error())
	}

	log.Infow("startup", "package", "db", "status", "ok")
}
