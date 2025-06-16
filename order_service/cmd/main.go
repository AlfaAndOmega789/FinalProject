package main

import (
	"gorm.io/gorm"
	"log"
	"net/http"
	"order/db"
	"order/internal/order/entity"
	"order/internal/order/handler"
	"order/internal/order/repository"
	"order/internal/order/usecase"
	"order/routes"
)

func main() {
	dbConn := db.InitDB()
	runMigrations(dbConn)

	orderRepo := repository.NewOrderRepository(dbConn)
	orderUsecase := usecase.NewOrderUsecase(orderRepo)
	orderHandler := handler.NewOrderHandler(orderUsecase)

	router := routes.SetupRouter(orderHandler)
	log.Println("Сервер запущен на :8082")
	log.Fatal(http.ListenAndServe(":8082", router))
}

func runMigrations(db *gorm.DB) {
	err := db.AutoMigrate(&entity.Order{})
	if err != nil {
		log.Fatal("Ошибка миграции:", err)
	}
}
