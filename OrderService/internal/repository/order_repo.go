package repository

import (
	"Payment/OrderService/internal/domain"
	"fmt"
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
	fmt.Println("UpdateOrder 1")
	var order domain.Orders
	if err := r.Database.First(&order, orderId).Error; err != nil {
		return order, err
	}
	fmt.Println("UpdateOrder 2")
	if err := r.Database.Model(&order).Update("status", NewStatus).Error; err != nil {
		return order, err
	}
	fmt.Println("UpdateOrder 3")
	return order, nil
}
