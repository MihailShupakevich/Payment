package repository

import (
	"Payment/internal/domain"
	"errors"
	"gorm.io/gorm"
)

type OrderRepo interface {
	PostOrder(order domain.Orders) (domain.Orders, error)
}
type OrderRepoImpl struct {
	Database *gorm.DB
}

func (r *OrderRepoImpl) PostOrder(order domain.Orders) (domain.Orders, error) {
	err := r.Database.Create(&order)
	if err != nil {
		return domain.Orders{}, errors.New("error creating order")
	}
	return order, nil
}
