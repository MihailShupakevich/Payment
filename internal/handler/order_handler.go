package handler

import (
	"Payment/internal/domain"
	"Payment/internal/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

type OrderHandler struct {
	uc usecase.OrderUsecase
}

type OrderHandlerInterface interface {
	PostOrder(ctx *gin.Context)
}

func (o *OrderHandler) PostOrder(c *gin.Context) {
	order := new(domain.Orders)
	err := c.BindJSON(order)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
		})
	}
	newOrder, err := o.uc.PostOrder(*order)
	c.JSON(http.StatusOK, newOrder)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err})
	}
}
