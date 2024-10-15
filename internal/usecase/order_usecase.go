package usecase

import (
	"Payment/internal/domain"
	"Payment/internal/repository"
)

type OrderUsecase struct {
	OrderRepo repository.OrderRepo
}

type OrderUseCaseInterface interface {
	PostOrder(order domain.Orders) (domain.Orders, error)
}

func (o *OrderUsecase) PostOrder(order domain.Orders) (domain.Orders, error) {
	newOrder, err := o.OrderRepo.PostOrder(order)
	if err != nil {
		return domain.Orders{}, err
	}
	return newOrder, nil
}
