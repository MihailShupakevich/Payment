package repository

import (
	"Payment/internal/domain"
	"gorm.io/gorm"
)

type OrderRepo interface {
	PostOrder(order domain.Orders) (domain.Orders, error)
}
type OrderRepoImpl struct {
	Database *gorm.DB
}

func New(database *gorm.DB) OrderRepo {
	return &OrderRepoImpl{
		Database: database,
	}
}

func (r *OrderRepoImpl) PostOrder(order domain.Orders) (domain.Orders, error) {
	nOrder := r.Database.Create(&order)
	if nOrder.Error != nil {
		return domain.Orders{}, nOrder.Error
	}
	return order, nil
}
