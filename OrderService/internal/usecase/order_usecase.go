package usecase

import (
	"Payment/internal/domain"
	"Payment/internal/repository"
	"fmt"
)

type OrderUsecase struct {
	OrderRepo repository.OrderRepo
}

type OrderUseCaseInterface interface {
	PostOrder(order domain.Orders) (domain.Orders, error)
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
