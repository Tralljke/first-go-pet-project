package infrastructure

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMySQLConnection(dsn string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	return db
}
