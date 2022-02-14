package database

import (
	"dev-book/src/config"
	"dev-book/src/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(config.DBConnectionString))
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&models.User{})
	return db, nil
}
