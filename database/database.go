package database

import (
	"log"
	"lollock/lock"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

// Get initializes the SQLite database with GORM.
func Get(dbPath string, logger *log.Logger) *gorm.DB {
	if db != nil {
		return db
	}
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Fatal("Failed to connect to database:", err)
	}
	err = db.AutoMigrate(&lock.Lock{})
	if err != nil {
		log.Fatal("Failed to migrate database schema:", err)
	}
	return db
}
