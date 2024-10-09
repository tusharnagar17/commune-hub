package database

import (
	"log"

	"github.com/tusharnagar17/commune-hub/internal/models"
	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) {
	if err := db.AutoMigrate(&models.User{}, &models.Event{}); err != nil {
		log.Fatal("Failed to migrate the database:", err)
	}

	log.Panicln("Database migrated successfully!")
}
