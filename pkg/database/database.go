package database

import (
	"database/sql"
	"gorm.io/gorm"
	"log"

	"guessr.net/pkg/config"

	_ "github.com/lib/pq"
)

var (
	DB *sql.DB
)

type Database struct {
	*gorm.DB
}

func SetupConnection() error {
	var db = DB
	masterDSN, _ := config.DbConfiguration()

	db, err := sql.Open("postgres", masterDSN)
	if err != nil {
		log.Println("Db connection error")
		return err
	}

	DB = db

	return nil
}
