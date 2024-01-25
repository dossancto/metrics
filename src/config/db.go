package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func GetDatabase() *gorm.DB {
	dsn := os.Getenv("POSTGRES_CONNECTION_STRING")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db
}
