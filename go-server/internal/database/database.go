package database

import (
	"log"

	"github.com/tusharnagar17/commune-hub/internal/config"
	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase(cfg config.Config) {
	var err error
	connStr := cfg.ConnStr()

	DB, err = gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
}

