package main

import (
	"Payment/internal/db"
	"Payment/internal/handler"
	"Payment/internal/repository"
	"Payment/internal/usecase"
	"github.com/gin-gonic/gin"
)

func main() {
	databaseConnect, err := db.Db()
	if err != nil {
		panic(err)
	}

	orderRepo := repository.New(databaseConnect)
	orderUsecase := usecase.New(orderRepo)
	orderHandler := handler.New(*orderUsecase)

	router := gin.Default()

	orderGroup := router.Group("/orders")
	orderHandler.SetupRoutes(orderGroup)
	router.Run(":8080")
}
