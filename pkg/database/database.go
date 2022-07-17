package database

import (
	"log"

	"guessr.net/pkg/config"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/dbresolver"
)

var (
	DB    *gorm.DB
	err   error
	DBErr error
)

type Database struct {
	*gorm.DB
}

func SetupConnection() error {
	var db = DB
	masterDSN, replicaDSN := config.DbConfiguration()

	logMode := viper.GetBool("DB_LOG_MODE")
	debug := viper.GetBool("DEBUG")

	loglevel := logger.Silent
	if logMode {
		loglevel = logger.Info
	}

	db, err = gorm.Open(postgres.Open(masterDSN), &gorm.Config{
		Logger: logger.Default.LogMode(loglevel),
	})
	if !debug {
		err := db.Use(dbresolver.Register(dbresolver.Config{
			Replicas: []gorm.Dialector{
				postgres.Open(replicaDSN),
			},
			Policy: dbresolver.RandomPolicy{},
		}))
		if err != nil {
			return err
		}
	}

	if err != nil {
		DBErr = err
		log.Println("Db connection error")
		return err
	}

	DB = db
	err := db.AutoMigrate(User{}, Song{})
	if err != nil {
		return err
	}

	return nil
}
