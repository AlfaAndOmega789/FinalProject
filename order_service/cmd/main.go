package main

import (
	"log"
	"net/http"
	"order/db"
	"order/internal/order/handler"
	"order/internal/order/repository"
	"order/internal/order/usecase"
	"order/routes"
)

func main() {
	dbConn := db.InitDB()

	orderRepo := repository.NewOrderRepository(dbConn)
	orderUsecase := usecase.NewOrderUsecase(orderRepo)
	orderHandler := handler.NewOrderHandler(orderUsecase)

	router := routes.SetupRouter(orderHandler)
	log.Println("Сервер запущен на :8082")
	log.Fatal(http.ListenAndServe(":8082", router))
}
