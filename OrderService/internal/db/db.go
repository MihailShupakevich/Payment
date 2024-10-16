package db

import (
	"Payment/OrderService/internal/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Db() (*gorm.DB, error) {
	dbData := "host=db dbname=payment user=admin password=admin port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dbData), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate([]domain.Orders{})
	return db, nil
}
