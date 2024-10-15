package handler

import (
	"Payment/internal/domain"
	"Payment/internal/usecase"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type OrderHandler struct {
	uc usecase.OrderUsecase
}

type OrderHandlerInterface interface {
	PostOrder(ctx *gin.Context)
}

func New(ucO usecase.OrderUsecase) OrderHandler {
	return OrderHandler{uc: ucO}
}

func (o *OrderHandler) SetupRoutes(router *gin.RouterGroup) {
	router.POST("/post", o.PostOrder)
}

func (o *OrderHandler) PostOrder(c *gin.Context) {
	order := new(domain.Orders)
	c.BindJSON(&order)
	fmt.Println(order)
	newOrder, err := o.uc.PostOrder(*order)
	fmt.Println(newOrder)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err})
	}
	c.JSON(http.StatusOK, newOrder)
}
