package db

import (
	"Payment/internal/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Db() (*gorm.DB, error) {
	dbData := "host=localhost dbname=payment user=admin password=admin sslmode=disable"
	db, err := gorm.Open(postgres.Open(dbData), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate([]domain.Orders{})
	return db, nil
}
