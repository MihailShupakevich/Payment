package usecase

import (
	"Payment/OrderService/internal/domain"
	"Payment/OrderService/internal/repository"
	"fmt"
)

type OrderUsecase struct {
	OrderRepo repository.OrderRepo
}

type OrderUseCaseInterface interface {
	PostOrder(order domain.Orders) (domain.Orders, error)
	UpdateOrder(orderId int, NewStatus string) (domain.Orders, error)
}

func New(orderRepo repository.OrderRepo) *OrderUsecase {
	return &OrderUsecase{OrderRepo: orderRepo}
}

func (o *OrderUsecase) PostOrder(order domain.Orders) (domain.Orders, error) {
	fmt.Println("UC 1")
	newOrder, err := o.OrderRepo.PostOrder(order)
	if err != nil {
		return domain.Orders{}, err
	}
	return newOrder, nil
}

func (o *OrderUsecase) UpdateOrder(orderId int, NewStatus string) (domain.Orders, error) {
	fmt.Println("UC 2")
	response, err := o.OrderRepo.UpdateOrder(orderId, NewStatus)
	if err != nil {
		return domain.Orders{}, err
	}
	return response, nil
}
