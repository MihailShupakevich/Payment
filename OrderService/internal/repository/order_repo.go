package repository

import (
	"Payment/OrderService/internal/domain"
	"gorm.io/gorm"
)

type OrderRepo interface {
	PostOrder(order domain.Orders) (domain.Orders, error)
	UpdateOrder(orderId int, NewStatus string) (domain.Orders, error)
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

func (r *OrderRepoImpl) UpdateOrder(orderId int, NewStatus string) (domain.Orders, error) {
	var order domain.Orders
	r.Database.Model(&order).Where("id = ?", orderId).First(&order).Update("status", NewStatus)
	return order, nil

}
